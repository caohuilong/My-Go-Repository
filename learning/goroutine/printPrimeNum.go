package main

import (
	"fmt"
	"runtime"
)

//var  wg sync.WaitGroup

func func1() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(2)

	wg.Add(2)
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func printPrime(name string) {
	defer wg.Done()

	for i := 2; i < 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%s: %d\n", name, i)
			//time.Sleep(1)
		}
	}
	fmt.Printf("Completed %s\n", name)
	return
}
