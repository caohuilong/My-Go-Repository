package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
)

func main1() {
	wg.Add(2)

	go incCounter("A")
	go incCounter("B")

	wg.Wait()
	fmt.Printf("Final counter=%d\n", counter)
}

func incCounter(name string) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		value := counter

		runtime.Gosched()

		value++
		counter = value
		fmt.Printf("%s counter=%d\n", name, counter)
	}
}