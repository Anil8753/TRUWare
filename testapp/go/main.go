package main

import (
	"log"
	"net/http"
)

func main() {

	h := Handler{}

	http.Handle("/", http.FileServer(http.Dir("./www")))
	http.HandleFunc("/api/query", h.Query)
	http.HandleFunc("/api/submit_txn", h.SubmitTransaction)

	log.Fatal(http.ListenAndServe(":8081", nil))

	// ledger := Ledger{}
	// err := ledger.Commit("a", "b", 10)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// amount, err := ledger.Query("a")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(amount)

}

// const (
// 	// Wallet file path
// 	walletFilePath = "../../network/vars/profiles/vscode/wallets/admin.truware.com"
// 	// Connection profile file path
// 	cpFilePath = "../../network/vars/profiles/mychannel_connection_for_gosdk.json"

// 	// User name
// 	user = "Admin"

// 	// App wallet path
// 	walletAppPath = "./connection/wallets"
// 	// Wallet file name
// 	walletFile = "Admin.id"

// 	// App connection profile path
// 	cpAppPath = "./connection"
// 	// App connection file name
// 	cpAppFile = "connection.json"

// 	// Channel name
// 	channelName = "mychannel"
// 	// Chaincode name
// 	chaincodeName = "simple"
// )

/*
To run this app, make sure that one of the wallet files such as Admin.id from
vars/profiles/vscode/wallets directory is copied onto ./wallets directory,
then this example code will use the wallet file and connection file to make
connections to Fabric network
*/

// func prepareArtifacts() error {

// 	PrintSectionHeader("Creating wallet artifacs...")

// 	// Wallet
// 	src, _ := filepath.Abs(path.Join(walletFilePath, walletFile))
// 	dst, _ := filepath.Abs(walletAppPath)

// 	if _, err := os.Stat(dst); !os.IsNotExist(err) {
// 		fmt.Println("Wallet artifacts already exist")
// 		return nil
// 	}

// 	if err := os.MkdirAll(dst, fs.ModePerm); err != nil {
// 		return fmt.Errorf("failed to create wallet directory: %v", err)
// 	}

// 	dstFile := path.Join(dst, walletFile)
// 	fmt.Println(src)
// 	fmt.Println(dst)

// 	err := CopyFile(src, dstFile)
// 	if err != nil {
// 		return fmt.Errorf("failed to copy wallet: %v", err)
// 	}

// 	// Connection profile
// 	src, _ = filepath.Abs(cpFilePath)
// 	dst, _ = filepath.Abs(path.Join(cpAppPath, cpAppFile))

// 	err = CopyFile(src, dst)
// 	if err != nil {
// 		return fmt.Errorf("failed to copy connection profile: %v", err)
// 	}

// 	PrintSectionHeader("Artifacts created successfully")
// 	return nil
// }

// func getContract() (*gateway.Contract, error) {

// 	if err := prepareArtifacts(); err != nil {
// 		return nil, err
// 	}

// 	wallet, err := gateway.NewFileSystemWallet(walletAppPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create wallet: %v", err)
// 	}

// 	if !wallet.Exists(user) {
// 		return nil, errors.New("failed to get Admin from wallet")
// 	}

// 	cp, _ := filepath.Abs(path.Join(cpAppPath, cpAppFile))

// 	gw, err := gateway.Connect(
// 		gateway.WithConfig(config.FromFile(cp)),
// 		gateway.WithIdentity(wallet, user),
// 	)

// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect: %v", err)
// 	}

// 	if gw == nil {
// 		return nil, errors.New("failed to create gateway")
// 	}

// 	network, err := gw.GetNetwork(channelName)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get network: %v", err)
// 	}

// 	return network.GetContract(chaincodeName), nil
// }

// func useWalletGateway() {

// 	contract, err := getContract()
// 	if err != nil {
// 		fmt.Println("failed: ", err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(contract.Name())

// 	start := time.Now()

// 	PrintSectionHeader("Submit transaction")

// 	result, err := contract.SubmitTransaction("invoke", "a", "b", "100")
// 	if err != nil {
// 		fmt.Printf("Failed to commit transaction: %v", err)
// 	} else {
// 		fmt.Println("Commit is successful", result)
// 	}

// 	PrintSectionHeader("Query Ledger")

// 	result, err = contract.EvaluateTransaction("query", "a")
// 	if err != nil {
// 		fmt.Printf("Failed to query: %v", err)
// 	} else {
// 		fmt.Println("Query result", string(result))
// 	}

// 	fmt.Println("The time took is ", time.Now().Sub(start))
// }

// // func main() {
// // 	useWalletGateway()
// // }
