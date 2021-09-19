package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type WarehouseContract struct {
	contractapi.Contract
}

func (s *WarehouseContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	// Create wallet store if not exist
	bytes, _ := ctx.GetStub().GetState(WalletKey)
	if bytes == nil {

		ws := make(WalletStore)
		bytes, err := json.Marshal(ws)
		if err != nil {
			return err
		}

		ctx.GetStub().PutState(WalletKey, bytes)
	}

	return nil
}
