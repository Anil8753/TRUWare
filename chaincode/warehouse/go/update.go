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

func (s *WarehouseContract) UpdateAsset(
	ctx contractapi.TransactionContextInterface,
	id string,
	assetJSON string,
) error {

	identity, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	// Read existing
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	if asset == nil {
		return errors.New("the asset does not exist")
	}

	// New Asset
	if err := json.Unmarshal([]byte(assetJSON), &asset); err != nil {
		return fmt.Errorf("invalid create assetJSON string. \n%s", assetJSON)
	}

	if err := ValidateAssetData(asset); err != nil {
		return err
	}

	if asset.OwnerId != identity {
		return errors.New("unauthorized warehouse owner")
	}

	bytes, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(asset.Id, bytes)
}
