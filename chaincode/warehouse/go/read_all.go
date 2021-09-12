package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ReadAllAsset returns all assets stored in the world state with the connected user (warehouse owner).
func (s *WarehouseContract) ReadAllAsset(
	ctx contractapi.TransactionContextInterface,
) ([]Asset, error) {

	identity, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	query := fmt.Sprintf("{\"selector\":{\"ownerId\":\"%s\"}}", identity)
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
