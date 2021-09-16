package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *OrderContract) PrematureClose(
	ctx contractapi.TransactionContextInterface,
	id string,
) error {

	identity, mspId, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	if mspId != CustomerPlaceOrderMSP {
		return fmt.Errorf("unauthorized user mspId: %s", mspId)
	}

	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("order id not exist. id:: %s", id)
	}

	order := Order{}
	if err := json.Unmarshal(bytes, &order); err != nil {
		return fmt.Errorf("failed to parse world state data. %v", err)
	}

	if identity != order.CustomerId {
		return fmt.Errorf("unauthorized user identity: %s", identity)
	}

	order.Status = PrematureClosed

	tokenPenality := (order.Value * order.PanalityPremature / 100)
	fmt.Println("---------------------------------------------------------")
	fmt.Println("token panality because of premature closer", tokenPenality)
	fmt.Println("---------------------------------------------------------")

	bytes, err = json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to marshal order. %v", order)
	}

	if err := ctx.GetStub().PutState(order.Id, bytes); err != nil {
		return fmt.Errorf("could not update the order into world state")
	}

	return nil
}
