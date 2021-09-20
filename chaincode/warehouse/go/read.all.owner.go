package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ReadAllOwnerAssets returns all assets stored in the world state with the connected user (warehouse owner).
func (s *WarehouseContract) ReadAllOwnerAssets(
	ctx contractapi.TransactionContextInterface,
) ([]Asset, error) {

	identity, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != WHOwnerReadAllMSP {
		return nil, fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	query := fmt.Sprintf(`{ "selector" : { "ownerId": "%s", "type": "warehouse" }}`, identity)
	itr, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query on world state: %v", err)
	}

	defer itr.Close()

	assets := make([]Asset, 0)

	for itr.HasNext() {

		res, err := itr.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate the query results: %v", err)
		}

		asset := Asset{}
		if err := json.Unmarshal(res.Value, &asset); err != nil {
			return nil, fmt.Errorf("failed to parse the query results: %v", err)
		}

		assets = append(assets, asset)
	}

	return assets, nil
}
