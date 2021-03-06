# TRUWare
A platform for efficient and cost effective usage of Warehouses.
Permissioned blockchain is the heart of this platform. It is implemented using Hyperledger Fabic blockchain technology, used Hyperledger Fabric version is **2.3.x**.
<br/>
Read more about the [HyperLedger Fabric.](https://www.hyperledger.org/)

## Components
It consists of the following components
- Network
  <br/>
  [minifabric](https://github.com/hyperledger-labs/minifabric) is used to setup the network infrastructure. 
  Follow the steps mentioned in the [network](./network) section to build the network.
  Refer the ```spec.yaml``` file to understand the initial organizations involved in the netwotk.
  
- Chaincode
  <br/>
  [Chaincode](./chaincode) section describes the neccessary details to develop, deploy and upgrade the chaincodes. [Go](https://golang.org/) is the programming language used for chaincode (smart contract) development.
  
- Web Application
  <br/>
  It has 2 parts, one API server and second Web interface.
  API server implementaion resides under the [warehouse/api](app/warehouse/api) and [customer/api](./app/customer/api) directories for different nodes. [Go](https://golang.org/) is the programming language used for blockchain interaction and API development.
  <br/>
  web It serves frontend interface of the blockchain system. It is implemented using [Angular](https://angular.io/). Source code is available in [app](./app) directory
  
- Scripts
  <br/>
  [scripts](./scripts) is the directory contains all the required scripts to build and run the application.
  
## Test Application
  Fully functional test application is available in [testapp](./testapp/go/) directory. It is implemented in Go programing language.
  It provides the web interface for [abstore chaincode](https://github.com/hyperledger/fabric-samples/blob/main/chaincode/abstore/go/abstore.go)
  
