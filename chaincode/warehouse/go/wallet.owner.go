package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) ReadOwnerWallet(
	ctx contractapi.TransactionContextInterface,
) (*WalletEntry, error) {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get identity. %v", err)
	}

	// Read sender
	ws, err := ReadWalletStore(ctx)
	if err != nil {
		return nil, err
	}

	we := ws[identity]
	return &we, nil
}
