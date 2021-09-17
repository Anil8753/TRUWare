package customer

import (
	"fmt"

	"github.com/Anil8753/truware/app/server/connector"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	channel         = "mychannel"
	ccOrderName     = "order"
	ccWarehouseName = "warehouse"
)

type Handler struct {
	// ccWarehouse *gateway.Contract
	ccOrder *gateway.Contract
}

func GetHandler() (*Handler, error) {

	nw, err := connector.GetNetwork(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to get network '%s'. \n%v", channel, err)
	}

	sc1 := connector.GetContract(nw, ccWarehouseName)
	if sc1 == nil {
		return nil, fmt.Errorf("failed to get chaincode '%s'", ccWarehouseName)
	}

	// sc2 := connector.GetContract(nw, ccWarehouseName)
	// if sc2 == nil {
	// 	return nil, fmt.Errorf("failed to get chaincode '%s'", ccWarehouseName)
	// }

	return &Handler{
		ccOrder: sc1,
		// ccWarehouse: sc2,
	}, nil
}
