package main

import (
	"fmt"
	"time"
)

func main() {
	// 데이터 크기 설정
	const dataSize = 10000
	keys := make([]string, dataSize)
	slice := make([]string, 0, dataSize)
	myMap := make(map[string]string)

	// 데이터 생성
	for i := 0; i < dataSize; i++ {
		key := fmt.Sprintf("key%d", i)
		keys[i] = key
		slice = append(slice, key)
		myMap[key] = fmt.Sprintf("value%d", i)
	}

	// 키 검색 벤치마크
	searchKey := "key5000"

	start := time.Now()
	for _, key := range slice {
		if key == searchKey {
			break
		}
	}
	fmt.Println("Slice search duration:", time.Since(start))

	start = time.Now()
	_ = myMap[searchKey]
	fmt.Println("Map search duration:", time.Since(start))
}
