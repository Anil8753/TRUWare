package warehouse

import (
	"github.com/Anil8753/truware/app/api/connector"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Handler struct {
	contract *gateway.Contract
}

func GetHandler() *Handler {

	sc, err := connector.GetContract("mychannel", "warehouse")
	if err != nil {
		panic(err)
	}

	return &Handler{
		contract: sc,
	}
}
