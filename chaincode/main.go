package main

import (
	"log"

	"chaincode2/chaincode"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(&chaincode.Contract{})
	if err != nil {
		log.Panicf("error creating chaincode: %v", err)
	}
	if err = chaincode.Start(); err != nil {
		log.Panicf("error starting chaincode: %v", err)
	}
}
