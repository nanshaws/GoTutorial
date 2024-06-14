package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j)
		time.Sleep(time.Second) // 模拟耗时工作
		fmt.Println("Worker", id, "finished job", j)
		results <- j * 2 // 将结果发送到results channel
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动3个worker goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送作业到jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭jobs channel，表示没有更多的作业

	// 收集结果
	for a := 1; a <= numJobs; a++ {
		result := <-results // 从results channel接收结果
		fmt.Println("Result:", result)
	}
}
