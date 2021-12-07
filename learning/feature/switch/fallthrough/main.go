package main

import "fmt"

func test(v int) {
	switch {
	case v < 3:
		fmt.Println(3)
	case v == 5:
		fmt.Println(5)
		fallthrough
	case v < 4:
		fmt.Println(4)
	default:
		fmt.Println("default")
	}
}

func main() {
	test(5)
}
