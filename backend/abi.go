package main

import (
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func loadAbi(filepath string) (*abi.ABI, error) {
	collectionFactoryJson, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(collectionFactoryJson)))
	if err != nil {
		return nil, err
	}

	return &contractAbi, nil
}
