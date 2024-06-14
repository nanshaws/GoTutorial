package main

import (
	"fmt"
	"sync"
	"time"
)

// Job 表示一个待执行的任务
type Job struct {
	id     int
	result chan<- int // 用于接收任务结果的通道
}

// worker 是工作池中的工作协程
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done() // 通知WaitGroup当前协程已结束
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j.id)
		time.Sleep(time.Second) // 模拟任务执行时间
		fmt.Printf("Worker %d finished job %d\n", id, j.id)
		j.result <- j.id * 2 // 假设任务是对id进行加倍
	}
}

func main() {
	const numJobs = 5 // 总任务数
	const numWorkers = 3 // 工作协程数

	jobs := make(chan Job, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// 启动工作协程
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// 发送任务
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j, result: results}
	}
	close(jobs) // 关闭任务通道，表示没有更多任务

	// 等待所有工作协程结束
	go func() {
		wg.Wait()
		close(results) // 所有工作协程结束后，关闭结果通道
	}()

	// 收集结果
	for result := range results {
		fmt.Println("Result:", result)
	}
}
