package main

import (
	"fmt"
	"strconv"
)

func main() {

	v64, err := strconv.ParseUint("1", 16, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v64)
}
