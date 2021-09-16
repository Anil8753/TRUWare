package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetInvokerIdentity(ctx contractapi.TransactionContextInterface) (string, string, error) {

	id, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", "", err
	}

	mspId, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", "", err
	}
	return id, mspId, nil
}
