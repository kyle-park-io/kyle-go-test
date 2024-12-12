package between

import (
	"fmt"
	"runtime"
)

func B() {
	// Measure memory usage before creating maps
	var memBefore runtime.MemStats
	runtime.ReadMemStats(&memBefore)

	// Create map[ElementType]struct{}
	set := make(map[int]struct{})
	for i := 0; i < 1000; i++ {
		set[i] = struct{}{}
	}

	// Measure memory usage after creating the first map
	var memAfterSet runtime.MemStats
	runtime.ReadMemStats(&memAfterSet)

	// Create map[ElementType]int
	multiset := make(map[int]int)
	for i := 0; i < 1000; i++ {
		multiset[i] = i
	}

	// Measure memory usage after creating the second map
	var memAfterMultiset runtime.MemStats
	runtime.ReadMemStats(&memAfterMultiset)

	// Calculate memory usage differences
	fmt.Printf("Memory usage for map[int]struct{}: %d bytes\n",
		memAfterSet.HeapAlloc-memBefore.HeapAlloc)
	fmt.Printf("Memory usage for map[int]int: %d bytes\n",
		memAfterMultiset.HeapAlloc-memAfterSet.HeapAlloc)
}
