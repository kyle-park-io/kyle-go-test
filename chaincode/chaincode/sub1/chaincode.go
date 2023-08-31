package sub1

import (
	"chaincode2/chaincode"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	chaincode.Contract
}

func (c *Contract) Test(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("sub1")
	return nil
}
