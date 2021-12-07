package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock = sync.RWMutex{}
	wg   = sync.WaitGroup{}
)

func writeLock(s string) {
	defer wg.Done()
	lock.Lock()
	fmt.Printf("goroutine %s get write lock\n", s)
	time.Sleep(1 * time.Second)
	fmt.Printf("goroutine %s release write lock\n", s)
	lock.Unlock()
}

func readLock(s string) {
	defer wg.Done()
	lock.RLock()
	fmt.Printf("goroutine %s get read lock\n", s)
	time.Sleep(1 * time.Second)
	fmt.Printf("goroutine %s release read lock\n", s)
	lock.RUnlock()
}

func main() {
	wg.Add(4)

	go writeLock("A")
	go writeLock("B")

	go readLock("C")
	go readLock("D")

	wg.Wait()
}
