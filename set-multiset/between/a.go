package between

import (
	"fmt"
	"reflect"
	"unsafe"
)

func A() {
	// map[ElementType]struct{}
	set := make(map[int]struct{})
	set[1] = struct{}{}
	set[2] = struct{}{}
	set[3] = struct{}{}

	// map[ElementType]int
	multiset := make(map[int]int)
	multiset[1] = 1
	multiset[2] = 2
	multiset[3] = 3

	// Function to calculate the approximate memory usage
	printMapMemoryUsage("set", set)
	printMapMemoryUsage("multiset", multiset)
}

func printMapMemoryUsage(name string, m interface{}) {
	// Get the reflect.Value of the map
	v := reflect.ValueOf(m)
	// Get the pointer to the map's underlying data structure
	ptr := v.Pointer()
	// Get the size of the map in bytes
	size := unsafe.Sizeof(m)

	fmt.Printf("Map: %s\n", name)
	fmt.Printf("Address: %p\n", ptr)
	fmt.Printf("Approx. Size: %d bytes\n\n", size)
}
