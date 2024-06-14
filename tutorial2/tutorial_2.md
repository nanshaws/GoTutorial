# Go语言基础教学文档

## 1. 概述

本教学文档旨在介绍Go语言的基础语法和特性，通过编写一个简单的程序来演示如何定义变量、结构体、方法、接口以及使用切片、映射等数据结构。

## 2. 变量和数据类型

Go语言支持多种数据类型，包括布尔型、整型、浮点型、复数、字符串等。以下是在代码中定义的几种变量及其数据类型：

```go
var isStudent bool = true
var age int = 30
var height float64 = 175.5
var complexNum complex128 = cmplx.Sqrt(-4)
var name string = "John Doe"
```

使用`fmt.Printf`函数可以打印变量的值和类型。

## 3. 结构体和方法

结构体是Go语言中自定义数据类型的一种方式。在代码中定义了一个`Person`结构体，并为其添加了一个`Introduce`方法：

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func (p *Person) Introduce() {
    fmt.Printf("Hi, my name is %s %s and I am %d years old.\n", p.FirstName, p.LastName, p.Age)
}
```

使用结构体时，可以创建结构体实例并调用其方法。

## 4. 接口

接口是Go语言中一种定义对象行为的类型。在代码中定义了一个`Executable`接口，并创建了一个实现该接口的`Command`结构体：

```go
type Executable interface {
    Execute() error
}

type Command struct {
    // Command的结构体字段
}

func (c *Command) Execute() error {
    // 实现接口方法
}
```

接口定义了一组方法的集合，结构体通过实现这些方法来实现接口。

## 5. 切片和映射

切片是Go语言中动态数组的实现方式，可以方便地添加和访问元素。映射是一种关联数组，用于存储键值对。在代码中定义了切片和映射，并展示了如何添加和访问元素：

```go
var sliceOfNumbers []int
var mapOfNumbers map[string]int

// 添加元素到映射
mapOfNumbers = map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}

// 使用append方法添加元素到切片
sliceOfNumbers = append(sliceOfNumbers, 1, 2, 3)
```

## 6. 程序流程

在`main`函数中，我们按照顺序执行了以下操作：

- 声明并初始化各种类型的变量。
- 创建结构体实例并调用其方法。
- 创建切片并添加元素。
- 创建映射并添加键值对。
- 使用`fmt.Printf`打印不同类型的变量。
- 创建`Command`实例并执行其`Execute`方法，处理可能出现的错误。

## 7. 总结

本教学文档通过演示一个简单的Go语言程序，介绍了变量、数据类型、结构体、方法、接口、切片和映射等

