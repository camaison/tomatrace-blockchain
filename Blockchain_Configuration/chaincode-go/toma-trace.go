package main

import (
	"log"
	"toma-trace/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	tomaTraceChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating toma-trace chaincode: %s", err.Error())
	}

	if err := tomaTraceChaincode.Start(); err != nil {
		log.Panicf("Error starting toma-trace chaincode: %s", err.Error())
	}
}
