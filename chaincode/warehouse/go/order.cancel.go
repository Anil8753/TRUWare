package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) CancelOrder(
	ctx contractapi.TransactionContextInterface,
	orderId string,
	comments string,
) error {

	log.Println("CancelOrder is invoked")

	_, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != ECOwnersMSP {
		return fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	log.Printf("invoker mspId: %s \n", mspId)

	// check if order exist
	order, err := s.readOrder(ctx, orderId)
	if err != nil {
		return fmt.Errorf("order id not exist. id: %s \n error: %v", orderId, err)
	}

	// Check if warehouse exist with (order.WarehouseId)
	wh, err := s.ReadAsset(ctx, order.WarehouseId)
	if err != nil {
		return fmt.Errorf("failed to read the warehouse data from world state. id: %s \n error: %v", order.WarehouseId, err)
	}

	// update order
	order.Status = PrematureClosed
	order.Conmments = comments

	tokenPenality := (order.Value * order.PanalityPremature / 100)

	if err := s.TransferToken(ctx, wh.OwnerId, strconv.Itoa(int(tokenPenality)), order.Conmments); err != nil {
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
	wh.General.AllocatedArea -= order.Space

	bytesWH, err := json.Marshal(wh)
	if err != nil {
		return fmt.Errorf("failed to marshal warehouse. %v", wh)
	}

	if err := ctx.GetStub().PutState(wh.Id, bytesWH); err != nil {
		return fmt.Errorf("could not write the warehouse to world state")
	}

	log.Println("order cancelled successfully")

	return nil
}
