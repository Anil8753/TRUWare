package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type OrderContract struct {
	contractapi.Contract
}

func (s *OrderContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}
