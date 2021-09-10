package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetInvokerIdentity(ctx contractapi.TransactionContextInterface) (string, error) {

	id, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", err
	}

	cert, err := ctx.GetClientIdentity().GetX509Certificate()
	if err != nil {
		return "", err
	}

	fmt.Printf("cert.PublicKey: %v\n", cert.PublicKey)

	return id, nil
}
