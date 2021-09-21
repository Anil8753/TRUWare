package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) UpdateOwnerRegistration(
	ctx contractapi.TransactionContextInterface,
	jsonStr string,
) error {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	re := RegistrationEntry{}
	if err := json.Unmarshal([]byte(jsonStr), &re); err != nil {
		return fmt.Errorf("invalid update identity json string. \nerror: %v\ninput data: %s", err, jsonStr)
	}

	if err := validateIdentityData(&re); err != nil {
		return fmt.Errorf("identity data validation failed. %v", err)
	}

	rs, err := ReadRegistrationStore(ctx)
	if err != nil {
		return fmt.Errorf("failed to read identity store. %v", err)
	}

	rs[identity] = re

	// Read sender
	if err := WriteRegistrationStore(ctx, rs); err != nil {
		return fmt.Errorf("failed to read the identity store. %v", err)
	}

	return nil
}

func validateIdentityData(re *RegistrationEntry) error {

	if re.Name == "" {
		return errors.New("name cannot be empty")
	}

	if re.Address == "" {
		return errors.New("address cannot be empty")
	}

	if re.Contact == "" {
		return errors.New("contact details cannot be empty")
	}

	if re.GSTNumber == "" {
		return errors.New("GST number cannot be empty")
	}

	return nil
}
