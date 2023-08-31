package sub2

import (
	"chaincode2/chaincode"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	chaincode.Contract
}

func (c *Contract) Test(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("sub2")
	return nil
}
