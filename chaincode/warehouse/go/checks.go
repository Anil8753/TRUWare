package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// assetExists returns true when asset with given ID exists in world state
func (s *WarehouseContract) assetExists(
	ctx contractapi.TransactionContextInterface,
	id string,
) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}
