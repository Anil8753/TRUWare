package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	// Wallet file path
	walletFilePath = "../../network/vars/profiles/vscode/wallets/admin.truware.com"
	// Connection profile file path
	cpFilePath = "../../network/vars/profiles/mychannel_connection_for_gosdk.json"

	// User name
	user = "Admin"

	// App wallet path
	walletAppPath = "./connection/wallets"
	// Wallet file name
	walletFile = "Admin.id"

	// App connection profile path
	cpAppPath = "./connection"
	// App connection file name
	cpAppFile = "connection.json"

	// Channel name
	channelName = "mychannel"
	// Chaincode name
	chaincodeName = "simple"
)

type Ledger struct {
}

func (l Ledger) Query(party string) (int, error) {

	contract, err := l.getContract()
	if err != nil {
		return 0, err
	}

	result, err := contract.EvaluateTransaction("query", party)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(result))
}

func (l Ledger) Commit(party1 string, party2 string, amount int) error {

	contract, err := l.getContract()
	if err != nil {
		return err
	}

	_, err = contract.SubmitTransaction("invoke", party1, party2, strconv.Itoa(amount))

	if err != nil {
		return err
	}

	return nil
}

// Private functions
func (l Ledger) prepareArtifacts() error {

	PrintSectionHeader("Creating wallet artifacs...")

	// Wallet
	src, _ := filepath.Abs(path.Join(walletFilePath, walletFile))
	dst, _ := filepath.Abs(walletAppPath)

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		fmt.Println("Wallet artifacts already exist")
		return nil
	}

	if err := os.MkdirAll(dst, fs.ModePerm); err != nil {
		return fmt.Errorf("failed to create wallet directory: %v", err)
	}

	dstFile := path.Join(dst, walletFile)
	fmt.Println(src)
	fmt.Println(dst)

	err := CopyFile(src, dstFile)
	if err != nil {
		return fmt.Errorf("failed to copy wallet: %v", err)
	}

	// Connection profile
	src, _ = filepath.Abs(cpFilePath)
	dst, _ = filepath.Abs(path.Join(cpAppPath, cpAppFile))

	err = CopyFile(src, dst)
	if err != nil {
		return fmt.Errorf("failed to copy connection profile: %v", err)
	}

	PrintSectionHeader("Artifacts created successfully")
	return nil
}

func (l Ledger) getContract() (*gateway.Contract, error) {

	if err := l.prepareArtifacts(); err != nil {
		return nil, err
	}

	wallet, err := gateway.NewFileSystemWallet(walletAppPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}

	if !wallet.Exists(user) {
		return nil, errors.New("failed to get Admin from wallet")
	}

	cp, _ := filepath.Abs(path.Join(cpAppPath, cpAppFile))

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(cp)),
		gateway.WithIdentity(wallet, user),
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

	return network.GetContract(chaincodeName), nil
}
