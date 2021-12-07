package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				lock.Lock()
				count = count + 1
				lock.Unlock()
			}
		}()
	}
	wg.Wait()

	fmt.Printf("count=%d\n", count)
}
