package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// UpdateTotalArea updates the rate of the warehouse.
func (s *WarehouseContract) UpdateRate(
	ctx contractapi.TransactionContextInterface,
	id string,
	rate int,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	if rate < 1 {
		return fmt.Errorf("rate cannot be less than 1. Enterer rate: %d", rate)
	}

	asset.General.Rate = rate

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
