package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) UpdateRegistration(
	ctx contractapi.TransactionContextInterface,
	jsonStr string,
) error {

	_, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	re, err := s.ReadRegistration(ctx)
	if err != nil {
		return fmt.Errorf("registration data not found. %v", err)
	}

	if err := json.Unmarshal([]byte(jsonStr), &re); err != nil {
		return fmt.Errorf("invalid update identity json string. \nerror: %v\ninput data: %s", err, jsonStr)
	}

	if err := validateRegistrationData(re); err != nil {
		return fmt.Errorf("identity data validation failed. %v", err)
	}

	bytes, err := json.Marshal(re)
	if err != nil {
		return fmt.Errorf("failed to marshal the registration data. %v", err)
	}

	if err := ctx.GetStub().PutState(re.ID, bytes); err != nil {
		return fmt.Errorf("failed to update the data into world state. %v", err)
	}

	return nil
}
