package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *OrderContract) ReadAllWarehouses(
	ctx contractapi.TransactionContextInterface,
) ([]Warehouse, error) {

	_, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != CustomerPlaceOrderMSP {
		return nil, fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	query := `{ "selector": { "type": "warehouse" }}`

	itr, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query on world state: %v", err)
	}

	defer itr.Close()

	warehouses := make([]Warehouse, 0)

	for itr.HasNext() {

		res, err := itr.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate the query results: %v", err)
		}

		wh := Warehouse{}
		if err := json.Unmarshal(res.Value, &wh); err != nil {
			return nil, fmt.Errorf("failed to parse the query results: %v", err)
		}

		warehouses = append(warehouses, wh)
	}

	return warehouses, nil
}
