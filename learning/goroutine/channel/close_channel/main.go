package main

import "fmt"

func main() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	for i := range c {
		fmt.Println(i)
	}

	c2 := make(chan int, 10)
	c2 <- 1
	close(c2)
	i, ok := <-c2
	fmt.Printf("%d, %t\n", i, ok)
	i, ok = <-c2
	fmt.Printf("%d, %t\n", i, ok)
}
