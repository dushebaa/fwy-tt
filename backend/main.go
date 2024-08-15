package main

import (
	"context"
	"os"

	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	nodeURL := os.Getenv("NODE_URI")

	factoryContractAddress := os.Getenv("FACTORY_CONTRACT_ADDRESS")
	var factorySContractTopics []common.Hash
	factorySContractTopics = append(factorySContractTopics, common.HexToHash(EVENT_COLLECTION_CREATED_TOPIC))
	factorySContractEvents := []string{EVENT_COLLECTION_CREATED_NAME}

	factorySContract := SContract{
		sEvents: factorySContractEvents,
		sTopics: factorySContractTopics,
	}
	err = factorySContract.init(
		common.HexToAddress(factoryContractAddress),
		ABI_COLLECTION_FACTORY_PATH,
	)
	if err != nil {
		log.Fatal(err)
	}

	var sContracts []SContract
	sContracts = append(sContracts, factorySContract)
	config := ScannerConfig{
		NodeURL:   nodeURL,
		Contracts: &sContracts,
	}

	scanner, err := NewScanner(config)
	if err != nil {
		log.Fatal(err)
	}

	go listenHttp(scanner)

	ctx := context.Background()
	if err := scanner.Scan(ctx); err != nil {
		log.Fatal(err)
	}
}
