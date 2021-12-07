package main

import "fmt"

func main() {
loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break loop
			}
			fmt.Printf("%d, %d\n", i, j)
		}
	}
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("done")
}
