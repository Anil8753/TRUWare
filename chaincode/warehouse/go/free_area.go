package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// FreeArea releases the specified space.
func (s *WarehouseContract) FreeArea(
	ctx contractapi.TransactionContextInterface,
	id string,
	allocationId string,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	index := -1
	bookingArea := 0

	for _, allocation := range asset.Allocations {
		if allocation.Id == allocationId {
			bookingArea = allocation.Area
			break
		}
		index++
	}

	if index != -1 {
		return fmt.Errorf("booking ID not found: %s", allocationId)
	}

	asset.General.AllocatedArea -= bookingArea
	asset.Allocations = append(asset.Allocations[:index], asset.Allocations[index+1:]...)

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
