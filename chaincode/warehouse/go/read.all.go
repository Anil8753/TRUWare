package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ReadAllAssets returns all assets stored in the world state with limited details.
func (s *WarehouseContract) ReadAllAssets(
	ctx contractapi.TransactionContextInterface,
) ([]Asset, error) {

	_, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != ECOwnersMSP {
		return nil, fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	query := "{\"selector\":{\"type\":\"warehouse\"}}"
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
