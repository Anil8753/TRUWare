package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) ReadRegistration(
	ctx contractapi.TransactionContextInterface,
) (*RegistrationEntry, error) {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	queryFmt := `{
		"selector": {
			"type": "registration",
			"ownerId": "%s"
		}
	}`

	query := fmt.Sprintf(queryFmt, identity)
	itr, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get query result. error: %v", err)
	}

	defer itr.Close()

	if !itr.HasNext() {
		return nil, errors.New("registration entry not found")
	}

	res, err := itr.Next()
	if err != nil {
		return nil, fmt.Errorf("shim.StateQueryIteratorInterface.Next() failed. error: %v", err)
	}

	re := RegistrationEntry{}
	if err := json.Unmarshal(res.Value, &re); err != nil {
		return nil, fmt.Errorf("failed to parse the query results: %v", err)
	}

	return &re, nil
}
