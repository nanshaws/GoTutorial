package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second) // 等待一秒钟
		fmt.Println(name, "says hello", i)
	}
}

func main() {
	go sayHello("Alice") // 启动一个goroutine执行sayHello函数
	go sayHello("Bob")   // 启动另一个goroutine执行sayHello函数

	// 主goroutine等待，防止程序直接退出
	time.Sleep(time.Second * 6)
	fmt.Println("Main function finished")
}
