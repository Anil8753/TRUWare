package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) ReadAllWarehouseOrders(
	ctx contractapi.TransactionContextInterface,
) ([]Order, error) {

	identity, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != WHOwnerReadAllMSP {
		return nil, fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	query := fmt.Sprintf(`{ "selector":{ "warehouseOwnerId": "%s", "type": "order" }}`, identity)

	itr, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query on world state: %v", err)
	}

	defer itr.Close()

	orders := make([]Order, 0)

	for itr.HasNext() {

		res, err := itr.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate the query results: %v", err)
		}

		order := Order{}
		if err := json.Unmarshal(res.Value, &order); err != nil {
			return nil, fmt.Errorf("failed to parse the query results: %v", err)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
