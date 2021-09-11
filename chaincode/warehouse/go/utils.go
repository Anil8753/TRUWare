package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetInvokerIdentity(ctx contractapi.TransactionContextInterface) (string, error) {

	id, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", err
	}

	return id, nil
}
