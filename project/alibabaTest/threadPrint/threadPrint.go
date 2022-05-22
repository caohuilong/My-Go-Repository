package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	// 使用两个channel实现两个goroutine的顺序打印
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	// 启动打印循环
	ch1 <- struct{}{}

	wg.Add(2)
	go Print1(ch1, ch2)
	go Print2(ch1, ch2)
	wg.Wait()
	return
}

func Print1(ch1 chan struct{}, ch2 chan struct{}) {
	defer wg.Done()
	for i := 1; i <= 100; i = i + 2 {
		<-ch1
		fmt.Printf("Print1-%d\n", i)
		ch2 <- struct{}{}
	}

	return
}

func Print2(ch1 chan struct{}, ch2 chan struct{}) {
	defer wg.Done()
	for i := 2; i <= 100; i = i + 2 {
		<-ch2
		fmt.Printf("Print2-%d\n", i)
		ch1 <- struct{}{}
	}

	return
}
