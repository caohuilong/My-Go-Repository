package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个接收信号的通道
	sigChan := make(chan os.Signal, 1)

	// 监听SIGTERM信号
	signal.Notify(sigChan, syscall.SIGTERM)

	// 启动一个 goroutine 来等待信号
	go func() {
		// 等待接收信号
		sig := <-sigChan
		fmt.Println("接收到信号:", sig)

		// 在这里执行你想要的操作，比如清理资源、保存状态等

		// 退出程序
		os.Exit(0)
	}()

	// 在这里执行你的主要逻辑

	// 阻塞主 goroutine，等待程序退出
	select {}
}