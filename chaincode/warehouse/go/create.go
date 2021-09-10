package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateAsset issues a new asset to the world state with given details.
// {
// 	"id":"100",
// 	"status":0,
// 	"generalInfo":{
// 	   "org":"rwc",
// 	   "name":"Rail Warehouse Corporation",
// 	   "phone": "9988776655",
// 	   "email": "support.rwc.com",
// 	   "address":"312/23 Bhanwandi, India",
// 	   "details":"One of the best warehouse in north india.",
// 	   "rate":8,
// 	   "totalArea":8000,
// 	   "allocatedArea":0
// 	},
// 	"postion":{
// 	   "latitude":12.99544495,
// 	   "longitude":77.75932179288739
// 	}
//  }

func (s *WarehouseContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	assetJSON string,
) error {

	identity, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	asset := Asset{}
	if err := json.Unmarshal([]byte(assetJSON), &asset); err != nil {
		return fmt.Errorf("invalid create assetJSON string. \n%s", assetJSON)
	}

	exists, err := s.AssetExists(ctx, asset.Id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("the asset '%s' already exists", asset.Id)
	}

	if asset.Id == "" ||
		asset.Postion.Latitude == 0 ||
		asset.Postion.Longitude == 0 ||
		asset.General.Org == "" ||
		asset.General.Address == "" ||
		asset.General.Name == "" ||
		asset.General.TotalArea == 0 ||
		asset.General.Details == "" {

		return errors.New(`these are mandatory fields.
							asset.Id, 
							asset.Postion.Latitude,
							asset.Postion.Longitude,
							asset.General.Org,
							asset.General.Address,
							asset.General.Name,
							asset.General.TotalArea,
							asset.General.Details`)
	}

	if asset.General.TotalArea < asset.General.AllocatedArea {
		return fmt.Errorf(
			"occupied area cannot be more than total area. occupied area: %d, total area: %d",
			asset.General.TotalArea, asset.General.AllocatedArea,
		)
	}

	if asset.General.Rate < 1 {
		return fmt.Errorf("rate cannot be less than 1. Rate: %d", asset.General.Rate)
	}

	asset.OwnerId = identity

	bytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(asset.Id, bytes)
}
