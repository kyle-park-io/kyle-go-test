package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			fmt.Println("tick")
		}
	}()

	time.Sleep(5 * time.Second)
}
