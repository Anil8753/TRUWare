package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&WarehouseContract{})
	if err != nil {
		log.Panicf("Error creating warehouse asset chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting warehouse asset chaincode: %v", err)
	}
}
