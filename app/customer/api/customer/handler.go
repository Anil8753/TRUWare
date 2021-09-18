package customer

import (
	"fmt"

	"github.com/Anil8753/truware/app/server/connector"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	channel         = "mychannel"
	ccWarehouseName = "warehouse"
)

type Handler struct {
	ccWarehouse *gateway.Contract
}

func GetHandler() (*Handler, error) {

	nw, err := connector.GetNetwork(channel)
	if err != nil {
		return nil, fmt.Errorf("failed to get network '%s'. \n%v", channel, err)
	}

	sc := connector.GetContract(nw, ccWarehouseName)
	if sc == nil {
		return nil, fmt.Errorf("failed to get chaincode '%s'", ccWarehouseName)
	}

	return &Handler{
		ccWarehouse: sc,
	}, nil
}
