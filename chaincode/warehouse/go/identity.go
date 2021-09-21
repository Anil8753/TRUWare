package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type IdentityEntry struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	GSTNumber string `json:"gst"`
}

const (
	IdentityKey = "identity"
)

type IdentityStore map[string]IdentityEntry

func ReadIdentityStore(
	ctx contractapi.TransactionContextInterface,
) (IdentityStore, error) {

	bytes, err := ctx.GetStub().GetState(IdentityKey)
	if err != nil {
		return nil, err
	}

	is := make(IdentityStore)
	if err := json.Unmarshal(bytes, &is); err != nil {
		return nil, err
	}

	return is, nil
}

func WriteIdentityStore(
	ctx contractapi.TransactionContextInterface,
	is IdentityStore,
) error {

	bytes, err := json.Marshal(is)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(IdentityKey, bytes); err != nil {
		return err
	}

	return nil
}
