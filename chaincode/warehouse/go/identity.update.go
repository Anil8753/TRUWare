package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) UpdateOwnerIdentity(
	ctx contractapi.TransactionContextInterface,
	jsonStr string,
) error {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	ie := IdentityEntry{}
	if err := json.Unmarshal([]byte(jsonStr), &ie); err != nil {
		return fmt.Errorf("invalid update identity json string. \nerror: %v\ninput data: %s", err, jsonStr)
	}

	if err := validateIdentityData(&ie); err != nil {
		return fmt.Errorf("identity data validation failed. %v", err)
	}

	is, err := ReadIdentityStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to read identity store. %v", err)
	}

	is[identity] = ie

	// Read sender
	if err := WriteIdentityStore(ctx, is); err != nil {
		return fmt.Errorf("failed to read the identity store. %v", err)
	}

	return nil
}

func validateIdentityData(ie *IdentityEntry) error {

	if ie.Name == "" {
		return errors.New("name cannot be empty")
	}

	if ie.Address == "" {
		return errors.New("address cannot be empty")
	}

	if ie.Contact == "" {
		return errors.New("contact details cannot be empty")
	}

	if ie.GSTNumber == "" {
		return errors.New("GST number cannot be empty")
	}

	return nil
}
