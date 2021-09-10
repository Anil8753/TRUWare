package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// UpdateGeneralDetails updates general details
func (s *WarehouseContract) UpdateGeneralDetails(
	ctx contractapi.TransactionContextInterface,
	id string,
	org string,
	name string,
	address string,
) error {

	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset.General.Org = org
	asset.General.Name = name
	asset.General.Address = address

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}
