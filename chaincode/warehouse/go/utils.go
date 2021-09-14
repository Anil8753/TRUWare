package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GetInvokerIdentity returns the current user's unique id and mspID
func GetInvokerIdentity(ctx contractapi.TransactionContextInterface) (string, string, error) {

	cid := ctx.GetClientIdentity()

	id, err := cid.GetID()
	if err != nil {
		return "", "", err
	}

	mspId, err := cid.GetMSPID()
	if err != nil {
		return "", "", err
	}

	return id, mspId, nil
}
