package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool, 1)

	// 启动一个goroutine来发送消息
	go func() {
		time.Sleep(time.Second*6)
		messages <- "ping"
	}()

	// 使用select语句等待两个channel中的一个可用
	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal:", sig)
	case <-time.After(7 * time.Second):
		fmt.Println("Timed out")
	}
}
