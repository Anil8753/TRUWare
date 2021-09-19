package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) BuyTokens(
	ctx contractapi.TransactionContextInterface,
	amount string,
	refNo string,
) error {

	log.Println("BuyTokens invoked")

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	amt, err := strconv.Atoi(amount)
	if err != nil {
		return fmt.Errorf("invaid amount entry: %s", amount)
	}

	ws, err := ReadWalletStore(ctx)
	if err != nil {
		return err
	}

	we, ok := ws[identity]
	// Create entry if not exist
	if !ok {
		we = WalletEntry{}
	}

	we.Owner = identity
	we.RefNo = refNo
	we.Balance = we.Balance + amt
	ws[identity] = we

	if err := WriteWalletStore(ctx, ws); err != nil {
		return err
	}

	return nil
}
