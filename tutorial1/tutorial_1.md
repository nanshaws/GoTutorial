# Go的安装

Go官网：[All releases - The Go Programming Language (google.cn)](https://golang.google.cn/dl/)

![image-20240610172827786](./../img/image-20240610172827786.png)

下载对应的版本，比如说Windows，就直接下载那个msi安装程序包就行了

# Hello World的简单逻辑

代码：

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
```

这个package  main  它告诉 Go 编译器这个文件包含可执行程序的入口点。Go 程序的执行总是从 `main` 包开始，而 `main` 函数是程序的入口点。因此，当你声明 `package main` 并且包含一个 `main` 函数时，Go 编译器会识别这个文件作为程序的启动文件。

打开cmd，执行以下命令

```
go mod init hello
```

Go 语言的模块初始化命令，用于创建一个新的 Go 模块并初始化其 `go.mod` 文件。这个命令是 Go 模块系统的一部分，它允许你管理项目的依赖项。

![image-20240610173602626](./../img/image-20240610173602626.png)

一般什么没有就模块和go的版本号

运行该程序，执行以下命令 

```
go run .
```

![image-20240610173718581](./../img/image-20240610173718581.png)

想编译成可执行文件，执行以下命令

```
go build .
```

![image-20240610173831192](./../img/image-20240610173831192.png)

# Go的基础语法

1. **包声明**： 每个 Go 程序都包含包（package），并且每个文件都属于一个包。Go 程序的入口点是 `main` 包中的 `main` 函数。

   ```
   package main
   ```

2. **导入语句**： 使用 `import` 关键字来导入其他包。

   ```
   import "fmt"
   import "os"
   ```

3. **变量声明**： 使用 `var` 关键字声明变量，变量的类型在变量名之后。

   ```
   var a int = 10
   ```

   或者使用短变量声明，这在 Go 中非常常见。

   ```
   var b = 20
   ```

4. **常量**： 使用 `const` 关键字定义常量。

   ```
   const Pi = 3.14159
   ```

5. **基本数据类型**： Go 有几种基本数据类型，包括整型（int、uint）、浮点型（float32、float64）、布尔型（bool）、字符串（string）等。

6. **控制结构**：

   - 条件语句：`if`、`else`。
   - 循环语句：`for`（Go 没有 `while` 或 `do-while` 循环，但可以用 `for` 实现相同的功能）。

   ```
   if x > 10 {
       fmt.Println("x 大于 10")
   } else {
       fmt.Println("x 小于或等于 10")
   }
   
   for i := 0; i < 10; i++ {
       fmt.Println(i)
   }
   ```

7. **函数**： 使用 `func` 关键字定义函数。

   ```
   func add(x int, y int) int {
       return x + y
   }
   ```

8. **数组和切片**： Go 中有固定长度的数组和动态的切片（slice）。

   ```
   var arr [5]int
   var slice []int = make([]int, 0)
   ```

9. **结构体**： 使用 `struct` 关键字定义结构体。

   ```
   type Point struct {
       X int
       Y int
   }
   ```

10. **接口**： 使用 `interface` 关键字定义接口。

    ```
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    ```

11. **并发**： Go 支持并发编程，使用 goroutines 和 channels。

    ```
    go someFunction()
    
    ch := make(chan int)
    ch <- 1
    ```

12. **错误处理**： Go 中的错误处理通常通过返回值来实现。

    ```
    f, err := os.Open("file.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
    ```

13. **包和模块**： Go 1.11 引入了模块系统，使用 `go mod` 命令管理依赖。

14. **格式化和注释**：

    - 使用 `//` 进行单行注释。
    - 使用 `/* ... */` 进行多行注释。
    - 使用 `fmt.Printf` 进行格式化输出。

15. **类型转换**： 显式类型转换使用类型名后跟圆括号。

    ```
    var x float64 = 3.0
    var y int = int(x)
    ```
