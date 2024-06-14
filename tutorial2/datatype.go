package main

import (
	"errors"
	"fmt"
	"math/cmplx"
)

// 定义一个结构体
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 给结构体添加方法
func (p *Person) Introduce() {
	fmt.Printf("Hi, my name is %s %s and I am %d years old.\n", p.FirstName, p.LastName, p.Age)
}

// 定义一个接口，描述了一个可以执行的操作
type Executable interface {
	Execute() error
}

// 定义一个函数，模拟执行操作并返回错误
func executeOperation() error {
	// 假设这里有一些操作，如果失败则返回错误
	// 这里我们返回一个示例错误
	return errors.New("operation failed")
}

// 定义一个结构体，它实现了Executable接口
type Command struct {
	// Command的结构体字段
}

// 实现Executable接口的Execute方法
func (c *Command) Execute() error {
	// 使用executeOperation函数执行操作，并处理错误
	err := executeOperation()
	if err != nil {
		// 错误处理：记录错误、重试、或者返回错误
		fmt.Println("Error occurred:", err)
		return err
	}
	// 如果没有错误，执行成功
	fmt.Println("Operation executed successfully")
	return nil
}

func main() {
	// 布尔型
	var isStudent bool = true

	// 整型
	var age int = 30

	// 浮点型
	var height float64 = 175.5

	// 复数
	var complexNum complex128 = cmplx.Sqrt(-4) // 90°的复数表示

	// 字符串
	var name string = "John Doe"

	// 数组
	var numbers [5]int

	// 切片
	var sliceOfNumbers []int

	// 映射
	var mapOfNumbers map[string]int

	// 声明并初始化结构体变量
	var person Person = Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       25,
	}

	// 给映射添加元素
	mapOfNumbers = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	// 打印不同类型的变量
	fmt.Printf("Boolean: %v\n", isStudent)
	fmt.Printf("Integer: %v\n", age)
	fmt.Printf("Float: %v\n", height)
	fmt.Printf("Complex Number: %v\n", complexNum)
	fmt.Printf("String: %v\n", name)
	fmt.Printf("Array: %v\n", numbers)
	fmt.Printf("Slice: %v\n", sliceOfNumbers)
	fmt.Printf("Map: %v\n", mapOfNumbers)

	// 调用结构体的方法
	person.Introduce()

	// 使用切片的 append 方法添加元素
	sliceOfNumbers = append(sliceOfNumbers, 1, 2, 3)
	fmt.Printf("Slice after append: %v\n", sliceOfNumbers)

	// 打印结构体变量
	fmt.Printf("Person: %+v\n", person)
	
	
	// 创建Command实例
	command := &Command{}

	// 执行Command的Execute方法，检查错误
	if err := command.Execute(); err != nil {
		fmt.Println("Failed to execute command:", err)
	}
}