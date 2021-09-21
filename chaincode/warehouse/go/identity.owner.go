package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) ReadOwnerRegistration(
	ctx contractapi.TransactionContextInterface,
) (*RegistrationEntry, error) {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get registration. %v", err)
	}

	// Read sender
	rs, err := ReadRegistrationStore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to read the registration store. %v", err)
	}

	re := rs[identity]
	return &re, nil
}
