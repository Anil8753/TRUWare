package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) ReadOwnerIdentity(
	ctx contractapi.TransactionContextInterface,
) (*IdentityEntry, error) {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	// Read sender
	is, err := ReadIdentityStore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read the identity store. %v", err)
	}

	ie := is[identity]
	return &ie, nil
}
