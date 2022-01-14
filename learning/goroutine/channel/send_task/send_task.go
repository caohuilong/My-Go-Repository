package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func task(ch chan int) {
	for t := range ch {
		fmt.Println("go task = ", t, ", goroutine count = ", runtime.NumGoroutine())
		wg.Done()
	}
}

func sendTask(t int, ch chan int) {
	wg.Add(1)
	ch <- t
}

func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go task(ch)
	}

	for i := 0; i < 10000; i++ {
		sendTask(i, ch)
	}

	wg.Wait()
}
