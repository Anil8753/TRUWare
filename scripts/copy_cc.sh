#!/bin/bash

rm -rf ../network/vars/chaincode/$1/go
cp -R ../chaincode/$1/go ../network/vars/chaincode/$1/go