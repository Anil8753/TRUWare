package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type RegistrationEntry struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Contact   string `json:"contact"`
	Email     string `json:"email"`
	GSTNumber string `json:"gst"`
}

const (
	RegistrationKey = "registration"
)

type IdentityStore map[string]RegistrationEntry

func ReadRegistrationStore(
	ctx contractapi.TransactionContextInterface,
) (IdentityStore, error) {

	bytes, err := ctx.GetStub().GetState(RegistrationKey)
	if err != nil {
		return nil, err
	}

	rs := make(IdentityStore)
	if err := json.Unmarshal(bytes, &rs); err != nil {
		return nil, err
	}

	return rs, nil
}

func WriteRegistrationStore(
	ctx contractapi.TransactionContextInterface,
	rs IdentityStore,
) error {

	bytes, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	if err := ctx.GetStub().PutState(RegistrationKey, bytes); err != nil {
		return err
	}

	return nil
}
