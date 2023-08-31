package main

import (
	"fmt"
	"math/big"
)

func main() {

	set := big.NewInt(0)
	_, ok := set.SetString("151623198801.18241914981489149891848914891894", 10)
	if !ok {
		fmt.Println("Failed to set big integer value")
		return
	}

	fmt.Println(set)
}
