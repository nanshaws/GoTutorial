# Go语言并发编程基础教学文档

## 1. 并发基础与Goroutines

在Go语言中，并发执行通常通过goroutine来实现。Goroutine是Go运行时环境中的轻量级线程，由Go运行时（runtime）管理，比操作系统线程更小，成千上万的goroutine可能同时存在于一个程序中。

### 示例1：简单的goroutine

```go
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
```

在上面的代码中，我们定义了一个`sayHello`函数，它会在一个循环中打印出问候语。在`main`函数中，我们使用`go`关键字启动了两个goroutine，分别执行`sayHello`函数并传入不同的名字。主goroutine使用`time.Sleep`等待一段时间，以确保其他goroutine有足够的时间执行完毕。

## 2. 通道（Channels）与并发通信

通道是Go语言中用于goroutine之间通信的内置原语。通道可以是有缓冲的或无缓冲的，它们用于在goroutine之间安全地传递值。

### 示例2：使用通道进行工作分配与结果收集

```go
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
```

在这个例子中，我们定义了一个`worker`函数，它接收一个作业ID、一个作业通道和一个结果通道。`main`函数中创建了作业通道和结果通道，并启动了多个worker goroutine。然后，我们向作业通道发送了多个作业，并关闭了作业通道。最后，我们从结果通道接收每个worker完成作业的结果。

## 3. 使用`select`语句进行多路选择

`select`语句用于在多个通信操作中选择一个可执行的操作。它常用于实现超时、等待多个通道中的消息等场景。

### 导入所需包

首先，我们需要导入`fmt`和`time`包。`fmt`包用于格式化输出，`time`包提供了时间相关的功能。

```go
import (
	"fmt"
	"time"
)
```

### 创建channel

在Go语言中，channel是用于goroutine间通信的管道。我们可以使用`make`函数来创建channel。

```go
messages := make(chan string)    // 创建一个用于发送字符串消息的channel
signals := make(chan bool, 1)    // 创建一个带缓冲的channel，用于发送布尔信号，容量为1
```

### 启动goroutine发送消息

接下来，我们启动一个goroutine来模拟发送消息的过程。这个goroutine将在等待6秒后向`messages` channel发送字符串`"ping"`。

```go
go func() {
	time.Sleep(time.Second * 6)
	messages <- "ping"
}()
```

### 使用select语句等待channel可用

`select`语句用于在多个通信操作中选择一个可执行的操作。如果多个case同时就绪，Go语言会随机选择一个执行。

```go
select {
case msg := <-messages:
	fmt.Println("Received message:", msg)
case sig := <-signals:
	fmt.Println("Received signal:", sig)
case <-time.After(7 * time.Second):
	fmt.Println("Timed out")
}
```

- `case msg := <-messages:`：等待从`messages` channel接收消息，并将其赋值给变量`msg`。
- `case sig := <-signals:`：等待从`signals` channel接收信号，并将其赋值给变量`sig`。
- `case <-time.After(7 * time.Second):`：等待7秒的超时。`time.After`函数返回一个在指定时间后发送当前时间的channel。如果其他case在超时前没有就绪，则执行这个case。

### 运行程序

当运行这个程序时，将看到以下输出之一：

- 如果`messages` channel在7秒内接收到消息，则输出`Received message: ping`。
- 如果`signals` channel接收到信号（本例中未使用），则输出`Received signal: <接收到的信号值>`。
- 如果在7秒内没有任何channel接收到消息或信号，则输出`Timed out`。

### 示例3：使用`select`实现超时机制

```go
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
```

# Go语言中的工作池模式与通道（Channel）使用教学

## 一、前言

在Go语言中，通道（Channel）是一种重要的并发原语，它允许在不同的goroutine之间进行安全的通信。同时，工作池（Worker Pool）模式是一种常用的并发处理方式，通过将任务分配给一组固定的goroutine来并发地执行任务。本教学文档将介绍如何在Go语言中实现工作池模式，并利用通道来传递任务和结果。

## 二、任务定义与通道创建

首先，我们定义一个`Job`结构体来表示待执行的任务。该结构体包含一个`id`字段用于标识任务，以及一个`result`通道用于接收任务执行后的结果。

```go
type Job struct {
	id     int
	result chan<- int // 用于接收任务结果的通道
}
```

接下来，我们创建两个通道：`jobs`用于发送任务，`results`用于接收任务执行的结果。

```go
jobs := make(chan Job, numJobs)
results := make(chan int, numJobs)
```

## 三、工作协程的创建与任务处理

我们定义一个`worker`函数，该函数将作为工作协程来执行任务。每个工作协程将从`jobs`通道接收任务，执行完毕后将结果发送到`result`通道，并通知`sync.WaitGroup`该协程已完成。

```go
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done() // 通知WaitGroup当前协程已结束
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j.id)
		// 模拟任务执行时间
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j.id)
		// 假设任务是对id进行加倍
		j.result <- j.id * 2
	}
}
```

在`main`函数中，我们根据所需的`numWorkers`数量创建相应数量的工作协程，并为每个协程分配一个唯一的ID。

```go
var wg sync.WaitGroup
for w := 1; w <= numWorkers; w++ {
	wg.Add(1)
	go worker(w, jobs, &wg)
}
```

## 四、发送任务与关闭通道

接下来，我们向`jobs`通道发送指定数量的任务，并在发送完毕后关闭该通道。

```go
for j := 1; j <= numJobs; j++ {
	jobs <- Job{id: j, result: results}
}
close(jobs) // 关闭任务通道，表示没有更多任务
```

关闭通道后，所有接收该通道的goroutine将收到一个零值或通道的关闭通知，从而结束循环。

## 五、等待工作协程结束与收集结果

在发送完所有任务后，我们启动一个新的goroutine来等待所有工作协程结束，并在所有工作协程结束后关闭`results`通道。

```go
go func() {
	wg.Wait()
	close(results) // 所有工作协程结束后，关闭结果通道
}()
```

最后，我们使用一个for循环从`results`通道中读取并打印每个任务的结果。

```go
for result := range results {
	fmt.Println("Result:", result)
}
```

由于`results`通道在所有工作协程结束后被关闭，这个循环将在读取完所有结果后自然结束。

## 六、总结

通过以上步骤，我们成功实现了一个基于工作池模式的并发任务处理系统。通过使用通道，我们能够安全地在goroutine之间传递任务和结果，同时利用`sync.WaitGroup`来确保主goroutine在所有任务完成后才退出。这种模式在处理大量并发任务时非常有效，因为它限制了并发执行的任务数量，从而避免了系统资源的过度消耗。

