# Go语言文件与目录操作教学文档

## 一、HTTP服务器搭建

### 1. 导入相关包

```go
import (
	"fmt"
	"log"
	"net/http"
)
```

### 2. 定义请求处理函数

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
```

### 3. 搭建HTTP服务器

```go
func main() {
	// 注册请求处理函数
	http.HandleFunc("/", helloHandler)

	// 打印启动信息
	log.Println("Server starting on :8080")

	// 监听端口并启动服务器
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

## 二、SQLite数据库操作

### 1. 导入相关包

```go
import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)
```

### 2. 操作数据库

```go
func main() {
	// 打开数据库文件
	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	_, err = db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
```

## 三、JSON编码与解码

### 1. 导入相关包

```go
import (
	"encoding/json"
	"fmt"
	"log"
)
```

### 2. 定义结构体

```go
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
```

### 3. 编码与解码JSON

```go
func main() {
	// 创建一个User对象
	user := User{ID: 1, Name: "Bob"}

	// 编码为JSON字符串
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData)) // 输出: {"id":1,"name":"Bob"}

	// 解码JSON字符串
	var decodedUser User
	if err := json.Unmarshal(jsonData, &decodedUser); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded User: %+v\n", decodedUser) // 输出: Decoded User: {ID:1 Name:Bob}
}
```

## 四、文件操作

### 一、简介

在Go语言中，对文件和目录的操作是非常常见的任务。Go标准库提供了`os`包，用于与操作系统进行交互，包括创建、读取、写入文件，以及创建、删除目录等操作。本教学文档将基于一个示例程序，详细讲解如何在Go语言中进行文件和目录操作。

### 二、示例程序

首先，我们来看一个完整的示例程序，该程序演示了如何创建文件、写入内容、读取内容、检查文件存在性、列出目录内容、创建新目录以及删除文件和目录。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 后续代码将在这里添加
}
```

### 三、操作详解

#### 1. 创建文件并写入内容

使用`os.WriteFile`函数可以一次性创建文件并写入内容。

```go
content := []byte("Hello, File!")
err := os.WriteFile("example.txt", content, 0644)
if err != nil {
	log.Fatal(err)
}
```

#### 2. 读取文件内容

使用`os.ReadFile`函数可以读取整个文件的内容。

```go
data, err := os.ReadFile("example.txt")
if err != nil {
	log.Fatal(err)
}
fmt.Println("File contents:", string(data))
```

#### 3. 检查文件是否存在

使用`os.Stat`函数可以获取文件或目录的信息，进而判断文件是否存在。

```go
if _, err := os.Stat("example.txt"); os.IsNotExist(err) {
	fmt.Println("File does not exist")
} else if err != nil {
	log.Fatal(err)
} else {
	fmt.Println("File exists")
}
```

#### 4. 列出目录内容

使用`os.ReadDir`函数可以列出指定目录下的所有文件和子目录。

```go
files, err := os.ReadDir(".")
if err != nil {
	log.Fatal(err)
}
for _, file := range files {
	fmt.Println(file.Name())
}
```

#### 5. 创建新目录

使用`os.Mkdir`函数可以创建新的目录。

```go
err = os.Mkdir("newdir", 0755)
if err != nil {
	log.Fatal(err)
}
```

#### 6. 删除文件

使用`os.Remove`函数可以删除指定的文件。

```go
err = os.Remove("example.txt")
if err != nil {
	log.Fatal(err)
}
```

#### 7. 删除目录

同样使用`os.Remove`函数删除目录，但需要注意目录必须为空，否则删除会失败。

```go
err = os.Remove("newdir")
if err != nil {
	log.Fatal(err)
}
```

### 四、注意事项

- 在进行文件或目录操作时，应始终检查返回的错误，以确保操作成功执行。
- 权限设置（如`0644`和`0755`）是Unix-like系统的权限设置方式，用于控制文件的读写执行权限。
- 删除目录时，必须确保目录为空，否则操作会失败。如果需要删除非空目录，可以使用`os.RemoveAll`函数。

### 五、总结

通过本教学文档，我们学习了如何在Go语言中使用`os`包进行文件和目录操作。这些操作包括创建文件、写入内容、读取内容、检查文件存在性、列出目录内容、创建新目录以及删除文件和目录。
