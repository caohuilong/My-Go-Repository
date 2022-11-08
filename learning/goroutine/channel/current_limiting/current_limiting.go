package main

import (
	"fmt"
	"sync"
	"time"
)

// 如何并发100个任务，但是同一时间最多运行的10个任务（waitGroup + channel）
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	c := make(chan int, 10)

	for i := 0; i < 100; i++ {
		c <- 1
		go func(i int) {
			fmt.Println(i)
			time.Sleep(1)
			wg.Done()
			<-c
		}(i)
	}
	wg.Wait()
}
