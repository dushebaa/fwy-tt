package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ScannerConfig struct {
	NodeURL              string
	ContractAddresses    []common.Address
	CollectionFactoryAbi *abi.ABI
	CollectionTokenAbi   *abi.ABI
}

type LogCollectionCreated struct {
	Collection common.Address
	Name       string
	Symbol     string
}

type Scanner struct {
	config       ScannerConfig
	client       *ethclient.Client
	logs         chan types.Log
	subscription *ethereum.Subscription

	collectionCreatedLogs []LogCollectionCreated
}

func NewScanner(config ScannerConfig) (*Scanner, error) {
	client, err := ethclient.Dial(config.NodeURL)
	if err != nil {
		return nil, err
	}

	query := ethereum.FilterQuery{
		Addresses: config.ContractAddresses,
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	return &Scanner{
		config:                config,
		client:                client,
		logs:                  logs,
		subscription:          &sub,
		collectionCreatedLogs: []LogCollectionCreated{},
	}, nil
}

func (s *Scanner) OnLogsRecieved(vLog types.Log) error {
	collectionCreatedEventRaw, err := s.config.CollectionFactoryAbi.Unpack("CollectionCreated", vLog.Data)
	if err != nil {
		return err
	}

	// dirty way to convert unpacked abi into a struct. Definitely never use
	// something like this for actual tx scanners
	collectionCreatedEvent := LogCollectionCreated{
		Collection: collectionCreatedEventRaw[0].(common.Address),
		Name:       collectionCreatedEventRaw[1].(string),
		Symbol:     collectionCreatedEventRaw[2].(string),
	}

	s.collectionCreatedLogs = append(s.collectionCreatedLogs, collectionCreatedEvent)
	return nil
}

func (s *Scanner) Scan(ctx context.Context) error {
	fmt.Println("Scanner initialized...")
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-(*s.subscription).Err():
			return err
		case vLog := <-s.logs:
			s.OnLogsRecieved(vLog)
		}
	}
}
