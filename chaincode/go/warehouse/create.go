package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateAsset issues a new asset to the world state with given details.
func (s *WarehouseContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	org string,
	name string,
	address string,
	totalArea int,
	bookedArea int,
	rate int,
) error {

	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	if totalArea < bookedArea {
		return fmt.Errorf(
			"occupied area cannot be more than total area. occupied area: %d, total area: %d",
			bookedArea, totalArea,
		)
	}

	if rate < 1 {
		return fmt.Errorf("rate cannot be less than 1. Rate: %d", rate)
	}

	asset := Asset{
		OwnerID:    "123",
		ID:         id,
		Org:        org,
		Name:       name,
		Address:    address,
		TotalArea:  totalArea,
		BookedArea: bookedArea,
		Rate:       rate,
		Status:     Operational,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
