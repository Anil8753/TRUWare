package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateAsset issues a new asset to the world state with given details.
// {
// 	"id":"100",
// 	"status":0,
// 	"generalInfo":{
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
		return fmt.Errorf("invalid create assetJSON string. \nerror: %v\ninput data: %s", err, assetJSON)
	}

	exists, err := s.AssetExists(ctx, asset.Id)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("the asset '%s' already exists", asset.Id)
	}

	if err := ValidateAssetData(&asset); err != nil {
		return err
	}

	asset.OwnerId = identity

	// if asset.Allocations == nil {
	// 	asset.Allocations = make([]Allocation, 0)
	// }

	bytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(asset.Id, bytes)
}
