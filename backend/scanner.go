package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ScannerConfig struct {
	NodeURL   string
	Contracts *[]SContract
}

type LogCollectionCreated struct {
	Collection common.Address
	Name       string
	Symbol     string
}

type LogTokenMinted struct {
	Collection common.Address
	Recipient  common.Address
	TokenId    big.Int
	TokenUri   string
}

type Scanner struct {
	config        ScannerConfig
	client        *ethclient.Client
	logs          chan types.Log
	subscriptions *[]Subscription
	wg            sync.WaitGroup
	ctx           context.Context

	collectionCreatedLogs []LogCollectionCreated
	tokensMintedLogs      map[string][]LogTokenMinted
}

func (s *Scanner) CreateSubscriptions(sContracts []SContract) (*[]Subscription, error) {
	subs := []Subscription{}
	for _, sContract := range sContracts {
		sContractAddresses := []common.Address{sContract.ContractAddress}

		topicsFilter := [][]common.Hash{sContract.sTopics}
		query := ethereum.FilterQuery{
			Addresses: sContractAddresses,
			Topics:    topicsFilter,
		}

		logs := make(chan types.Log)
		sub, err := s.client.SubscribeFilterLogs(context.Background(), query, logs)
		if err != nil {
			return nil, err
		}

		for _, sEvent := range sContract.sEvents {
			sSub := Subscription{
				sSub:       sub,
				sLogs:      &logs,
				sContract:  &sContract,
				sEventName: sEvent,
			}

			subs = append(subs, sSub)
		}
	}

	return &subs, nil
}

func NewScanner(config ScannerConfig) (*Scanner, error) {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		return nil, err
	}

	sc := &Scanner{
		config: config,
		client: client,
		wg:     sync.WaitGroup{},

		tokensMintedLogs: make(map[string][]LogTokenMinted),
	}

	subs, err := sc.CreateSubscriptions(*config.Contracts)
	if err != nil {
		return nil, err
	}

	sc.subscriptions = subs
	return sc, nil
}

func (s *Scanner) OnLogsReceived(vLog types.Log, sub Subscription) error {
	fmt.Println("Received event:", sub.sEventName)
	txLogsRaw, err := sub.sContract.ContractAbi.Unpack(sub.sEventName, vLog.Data)
	if err != nil {
		return err
	}

	switch sub.sEventName {
	case EVENT_COLLECTION_CREATED_NAME:
		collectionCreatedEvent := LogCollectionCreated{
			Collection: txLogsRaw[0].(common.Address),
			Name:       txLogsRaw[1].(string),
			Symbol:     txLogsRaw[2].(string),
		}
		s.collectionCreatedLogs = append(s.collectionCreatedLogs, collectionCreatedEvent)

		collectionSContract := SContract{
			sEvents: []string{EVENT_TOKEN_MINTED_NAME},
			sTopics: []common.Hash{common.HexToHash(EVENT_TOKEN_MINTED_TOPIC)},
		}
		err = collectionSContract.init(
			txLogsRaw[0].(common.Address),
			ABI_COLLECTION_TOKEN_PATH,
		)
		if err != nil {
			return err
		}
		sub, err := s.CreateSubscriptions([]SContract{collectionSContract})
		if err != nil {
			return err
		}
		// subscribe to collection nft mint events
		s.AddSubscription((*sub)[0])
	case EVENT_TOKEN_MINTED_NAME:
		tokenAddress := txLogsRaw[0].(common.Address)
		tokenMintedEvent := LogTokenMinted{
			Collection: tokenAddress,
			Recipient:  txLogsRaw[1].(common.Address),
			TokenId:    *txLogsRaw[2].(*big.Int),
			TokenUri:   txLogsRaw[3].(string),
		}
		tokenAddressLowercase := strings.ToLower(tokenAddress.Hex())
		s.tokensMintedLogs[tokenAddressLowercase] = append(s.tokensMintedLogs[tokenAddressLowercase], tokenMintedEvent)
	}

	return nil
}

func (s *Scanner) AddSubscription(sub Subscription) error {
	for {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		case err := <-sub.sSub.Err():
			return err
		case vLog := <-*sub.sLogs:
			s.OnLogsReceived(vLog, sub)
		}
	}
}

func (s *Scanner) Scan(ctx context.Context) error {
	fmt.Println("Scanner initialized...")

	s.ctx = ctx
	errChan := make(chan error, len(*s.subscriptions))

	for _, sub := range *s.subscriptions {
		s.wg.Add(1)
		go func(sub Subscription) {
			defer s.wg.Done()
			for {
				select {
				case <-s.ctx.Done():
					errChan <- s.ctx.Err()
					return
				case err := <-sub.sSub.Err():
					errChan <- err
					return
				case vLog := <-*sub.sLogs:
					s.OnLogsReceived(vLog, sub)
				}
			}
		}(sub)
	}

	go func() {
		s.wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil

}
