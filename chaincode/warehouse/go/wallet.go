package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type WalletEntry struct {
	Owner   string `json:"owner"`
	Balance int    `json:"balance"`
	RefNo   string `json:"refNo"`
}

const (
	WalletKey = "wallet"
)

type WalletStore map[string]WalletEntry

func ReadWalletStore(
	ctx contractapi.TransactionContextInterface,
) (WalletStore, error) {

	bytes, err := ctx.GetStub().GetState(WalletKey)
	if err != nil {
		return nil, err
	}

	ws := make(WalletStore)
	if err := json.Unmarshal(bytes, &ws); err != nil {
		return nil, err
	}

	return ws, nil
}

func WriteWalletStore(
	ctx contractapi.TransactionContextInterface,
	ws WalletStore,
) error {

	bytes, err := json.Marshal(ws)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(WalletKey, bytes); err != nil {
		return err
	}

	return nil
}
