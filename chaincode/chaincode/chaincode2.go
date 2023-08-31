package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type A struct {
	name string
}

type Contract2 struct {
	contractapi.Contract
}

func (c *Contract2) Test(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("chaincode2")
	return nil
}
