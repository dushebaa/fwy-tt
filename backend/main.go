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
	var contractAddresses []common.Address

	factoryContractAddress := os.Getenv("FACTORY_CONTRACT_ADDRESS")
	contractAddresses = append(contractAddresses, common.HexToAddress(factoryContractAddress))

	factoryContractAbi, err := loadAbi("abi/collectionFactory.json")
	if err != nil {
		log.Fatal(err)
	}

	config := ScannerConfig{
		NodeURL:              nodeURL,
		ContractAddresses:    contractAddresses,
		CollectionFactoryAbi: factoryContractAbi,
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
