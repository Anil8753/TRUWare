package connector

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// GetNetwork returns the contract
func GetNetwork(channelName string) (*gateway.Network, error) {

	cfg := NewConfig()

	wallet, err := gateway.NewFileSystemWallet(cfg.WalletPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	if !wallet.Exists(cfg.User) {
		return nil, errors.New("failed to get Admin from wallet")
	}

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(cfg.ConnectionProfilePath)),
		gateway.WithIdentity(wallet, cfg.User),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	if gw == nil {
		return nil, errors.New("failed to create gateway")
	}

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		return nil, fmt.Errorf("failed to get network: %v", err)
	}

	return network, nil
}

// GetContract returns the contract
func GetContract(network *gateway.Network, contract string) *gateway.Contract {
	return network.GetContract(contract)
}
