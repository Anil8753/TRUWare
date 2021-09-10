package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// BookArea books the specified area from the total available space.
func (s *WarehouseContract) BookArea(
	ctx contractapi.TransactionContextInterface,
	id string,
	allocatedArea int,
	duration int,
) error {

	identity, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	available := asset.General.TotalArea - asset.General.AllocatedArea

	if (available) < allocatedArea {
		return fmt.Errorf(
			"not enough space. available area: %d,  requested booking area %d",
			available, allocatedArea,
		)
	}

	allocationId := strconv.Itoa(len(asset.Allocations) + 1)

	asset.Allocations = append(asset.Allocations, Allocation{
		Id:       allocationId,
		ClientId: identity,
		Area:     allocatedArea,
		Duration: duration,
	})

	asset.General.AllocatedArea += allocatedArea

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
