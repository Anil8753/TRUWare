package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric/common/util"
)

func (s *OrderContract) readWarehouse(
	ctx contractapi.TransactionContextInterface,
	id string,
) (*Warehouse, error) {

	fmt.Printf("trying to read the warehouse with id: %s \n", id)

	chainCodeArgs := util.ToChaincodeArgs("ReadAsset", id)
	response := ctx.GetStub().InvokeChaincode("warehouse", chainCodeArgs, "mychannel")

	if response.Status != shim.OK {
		errMsg := fmt.Errorf("could not found the warehouse. id: %s \n. error: %s", id, response.Message)
		return nil, errMsg
	}

	bytes := response.GetPayload()

	wh := Warehouse{}
	if err := json.Unmarshal(bytes, &wh); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data received from world state.\n %v \n error: %v", string(bytes), err)
	}

	fmt.Println("got the warehouse")
	{
		b, err := ctx.GetStub().GetState(id)
		if err != nil {
			fmt.Printf("could not found the warehouse. id: %s", id)
		}

		wh2 := Warehouse{}
		if err := json.Unmarshal(b, &wh2); err != nil {
			fmt.Printf("failed to unmarshal data received from world state.\n %v \n error: %v", string(b), err)
		}
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

	chainCodeArgs := util.ToChaincodeArgs("UpdateAsset", id, string(bytes))
	response := ctx.GetStub().InvokeChaincode("warehouse", chainCodeArgs, "mychannel")
	if response.Status != shim.OK {
		return fmt.Errorf("could not update the warehouse. id: %s \n. error: %s", id, response.Message)
	}

	// if err := ctx.GetStub().PutState(id, bytes); err != nil {
	// 	return fmt.Errorf("could not write the warehouse to world state")
	// }

	fmt.Printf("warehouse data updated. %s\n", string(bytes))
	log.Printf("warehouse data updated. %s\n", string(bytes))

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
