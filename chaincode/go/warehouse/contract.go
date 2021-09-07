package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type WarehouseContract struct {
	contractapi.Contract
}

func (s *WarehouseContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}
