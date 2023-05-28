package main

import "fmt"

func main() {
	testMap := make(map[int]struct{})
	testMap[1] = struct{}{}
	testMap[2] = struct{}{}

	tmp := testMap[2]
	if tmp != struct{}{} {
		fmt.Println("3 not exist")
	}

	// delete all element in the testMap
	for k := range testMap {
		delete(testMap, k)
	}

	return
}
