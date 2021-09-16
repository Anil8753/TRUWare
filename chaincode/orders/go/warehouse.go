package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *OrderContract) readWarehouse(
	ctx contractapi.TransactionContextInterface,
	id string,
) (*Warehouse, error) {

	bytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("could not found the warehouse. id: %s", id)
	}

	wh := Warehouse{}
	if err := json.Unmarshal(bytes, &wh); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data received from world state.\n %v", string(bytes))
	}

	return &wh, nil
}

func (s *OrderContract) updateWarehouse(
	ctx contractapi.TransactionContextInterface,
	id string,
	wh *Warehouse,
) error {

	bytes, err := json.Marshal(wh)
	if err != nil {
		return fmt.Errorf("failed to marshal warehouse. %v", wh)
	}

	if err := ctx.GetStub().PutState(id, bytes); err != nil {
		return fmt.Errorf("could not write the warehouse to world state")
	}

	return nil
}

type Warehouse struct {
	Type    string          `json:"type"`
	Id      string          `json:"id"`
	OwnerId string          `json:"ownerId"`
	General GeneralInfo     `json:"generalInfo"`
	Status  WarehouseStatus `json:"status"`
}

type WarehouseStatus int

const (
	Operational WarehouseStatus = iota
	NonOperational
)

type GeneralInfo struct {
	Name               string  `json:"name"`
	Phone              string  `json:"phone"`
	Email              string  `json:"email"`
	Address            string  `json:"address"`
	Details            string  `json:"details"`
	Rate               float64 `json:"rate"`
	PanalityAfterLimit float64 `json:"panalityAfterLimit"`
	PanalityPremature  float64 `json:"panalityPremature"`
	TotalArea          int     `json:"totalArea"`
	AllocatedArea      int     `json:"allocatedArea"`
}
