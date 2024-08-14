package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type SContract struct {
	ContractAbi     *abi.ABI
	ContractAddress common.Address

	sEvents []string
	sTopics []common.Hash
}

func (sc *SContract) init(contractAddress common.Address, abiPath string) error {
	contractAbi, err := loadAbi(abiPath)
	if err != nil {
		return err
	}

	sc.ContractAddress = contractAddress
	sc.ContractAbi = contractAbi

	return nil
}
