package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

func (c *Contract) Test(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("chaincode1")
	return nil
}
