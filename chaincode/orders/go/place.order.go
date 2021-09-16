package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *OrderContract) PlaceOrder(
	ctx contractapi.TransactionContextInterface,
	jsonStr string,
) error {

	identity, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != CustomerPlaceOrderMSP {
		return fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	order := Order{}
	if err := json.Unmarshal([]byte(jsonStr), &order); err != nil {
		return fmt.Errorf("invalid create assetJSON string. \nerror: %v\ninput data: %s", err, jsonStr)
	}

	if _, err := ctx.GetStub().GetState(order.Id); err == nil {
		return fmt.Errorf("order id already exist. id:: %s", order.Id)
	}

	// Check if warehouse exist with (order.WarehouseId)
	wh, err := s.readWarehouse(ctx, order.WarehouseId)
	if err != nil {
		return fmt.Errorf("failed to read the warehouse data from world state. id: %s", order.WarehouseId)
	}

	if wh.Status != Operational {
		return fmt.Errorf("warehouse is not operational.\n %v", wh.Status)
	}

	if order.Space <= 0 {
		return fmt.Errorf("invalid requested space: %d", order.Space)
	}

	if order.Duration <= 0 {
		return fmt.Errorf("invalid requested duration: %d", order.Space)
	}

	available := wh.General.TotalArea - wh.General.AllocatedArea
	if available < order.Space {
		return fmt.Errorf("not enough space.\n requested: %d. available: %d", order.Space, available)
	}

	// return &asset, nil
	order.Type = "order"
	order.Status = Placed
	order.CustomerId = identity
	order.WarehouseId = wh.Id
	order.Rate = wh.General.Rate
	order.Value = wh.General.Rate * float64(order.Space)
	order.PanalityAfterLimit = wh.General.PanalityAfterLimit
	order.PanalityPremature = wh.General.PanalityPremature

	// update warehouse allocated space
	wh.General.AllocatedArea += order.Space
	if err := s.updateWarehouse(ctx, wh.Id, wh); err != nil {
		return fmt.Errorf("failed to update the warehouse allocated space")
	}

	bytes, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal order. %v", order)
	}

	if err := ctx.GetStub().PutState(order.Id, bytes); err != nil {
		return fmt.Errorf("could not write the order to world state")
	}

	return nil
}
