package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// UpdateTotalArea updates the total area of the warehouse.
// It checks that total area should not be less than the occupied area.
func (s *WarehouseContract) UpdateTotalArea(
	ctx contractapi.TransactionContextInterface,
	id string,
	totalArea int,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	if asset.General.AllocatedArea < totalArea {
		return fmt.Errorf(
			"total area (%d) cannot be less than occupied area (%d)",
			totalArea, asset.General.AllocatedArea,
		)
	}

	asset.General.TotalArea = totalArea

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
