package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func (s *WarehouseContract) TransferToken(
	ctx contractapi.TransactionContextInterface,
	receiver string,
	amount string,
	refNo string,
) error {

	identity, _, err := GetInvokerIdentity(ctx)
	if err != nil {
		return fmt.Errorf("failed to get identity. %v", err)
	}

	ws, err := ReadWalletStore(ctx)
	if err != nil {
		return err
	}

	weSender, ok := ws[identity]
	if !ok {
		return fmt.Errorf("wallet does not exist, sender: %s", identity)
	}

	amt, err := strconv.Atoi(amount)
	if err != nil {
		return fmt.Errorf("invaid amount entry: %s", amount)
	}

	if weSender.Balance < amt {
		return fmt.Errorf("not enough balance: %s", amount)
	}

	weReceiver, ok := ws[receiver]
	if !ok {
		weReceiver = WalletEntry{}
	}

	weSender.Balance -= amt
	weSender.RefNo = refNo
	ws[identity] = weSender

	weReceiver.Balance += amt
	weReceiver.RefNo = refNo
	ws[receiver] = weReceiver

	return WriteWalletStore(ctx, ws)
}
