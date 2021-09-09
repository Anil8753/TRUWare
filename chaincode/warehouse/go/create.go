package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// CreateAsset issues a new asset to the world state with given details.
// {
// 	"OwnerID": "123",
// 	"ID": "xyz",
// 	"Org": "warebox",
// 	"Name": "warebox inc",
// 	"Address": "212/13 Ganga Nagar, Kengeri",
// 	"TotalArea": 5000,
// 	"BookedArea": 0,
// 	"Rate": 100,
// 	"Postion": {
// 	  "Latitude": 100,
// 	  "Longitude": 290.67
// 	},
// 	"Status": 0,
// 	"Bookings": []
// }
func (s *WarehouseContract) CreateAsset(
	ctx contractapi.TransactionContextInterface,
	assetJSON string,
) error {

	asset := Asset{}
	if err := json.Unmarshal([]byte(assetJSON), &asset); err != nil {
		return fmt.Errorf("invalid create assetJSON string. \n%s", assetJSON)
	}

	exists, err := s.AssetExists(ctx, asset.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset '%s' already exists", asset.ID)
	}

	if asset.TotalArea < asset.BookedArea {
		return fmt.Errorf(
			"occupied area cannot be more than total area. occupied area: %d, total area: %d",
			asset.TotalArea, asset.BookedArea,
		)
	}

	if asset.Rate < 1 {
		return fmt.Errorf("rate cannot be less than 1. Rate: %d", asset.Rate)
	}

	return ctx.GetStub().PutState(asset.ID, []byte(assetJSON))
}
