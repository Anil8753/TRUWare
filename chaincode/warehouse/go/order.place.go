package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) PlaceOrder(
	ctx contractapi.TransactionContextInterface,
	jsonStr string,
) error {

	log.Println("PlaceOrder is invoked")

	identity, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != ECOwnersMSP {
		return fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	log.Printf("invoker mspId: %s \n", mspId)

	order := Order{}
	if err := json.Unmarshal([]byte(jsonStr), &order); err != nil {
		return fmt.Errorf("invalid create assetJSON string. \nerror: %v\ninput data: %s", err, jsonStr)
	}

	// check if order is already exist
	if bytes, _ := ctx.GetStub().GetState(order.Id); bytes != nil {
		return fmt.Errorf("order id already exist. id:: %s", order.Id)
	}

	// Check if warehouse exist with (order.WarehouseId)
	wh, err := s.ReadAsset(ctx, order.WarehouseId)
	if err != nil {
		return fmt.Errorf("failed to read the warehouse data from world state. id: %s \n error: %v", order.WarehouseId, err)
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

	// update order
	order.Type = "order"
	order.Status = Placed
	order.CustomerId = identity
	order.WarehouseId = wh.Id
	order.Rate = wh.General.Rate
	order.Value = wh.General.Rate * float64(order.Space)
	order.PanalityAfterLimit = wh.General.PanalityAfterLimit
	order.PanalityPremature = wh.General.PanalityPremature

	// Transfer token
	if err := s.TransferToken(ctx, wh.OwnerId, strconv.Itoa(int(order.Value)), order.Conmments); err != nil {
		return fmt.Errorf("failed to transfer token. %v", err)
	}

	bytes, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal order. %v", order)
	}

	if err := ctx.GetStub().PutState(order.Id, bytes); err != nil {
		return fmt.Errorf("could not write the order to world state")
	}

	// update warehouse allocated space
	wh.General.AllocatedArea += order.Space

	bytesWH, err := json.Marshal(wh)
	if err != nil {
		return fmt.Errorf("failed to marshal warehouse. %v", wh)
	}

	if err := ctx.GetStub().PutState(wh.Id, bytesWH); err != nil {
		return fmt.Errorf("could not write the warehouse to world state")
	}

	fmt.Println("order placed successfully")

	return nil
}
