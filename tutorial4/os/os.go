package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建一个新文件并写入内容
	content := []byte("Hello, File!")
	err := os.WriteFile("example.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// 读取文件内容
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File contents:", string(data))

	// 检查文件是否存在
	if _, err := os.Stat("example.txt"); os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File exists")
	}

	// 列出目录内容
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// 创建新目录
	err = os.Mkdir("newdir", 0755)
	if err != nil {
		log.Fatal(err)
	}

	// 删除文件
	err = os.Remove("example.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 删除目录（注意：目录必须为空）
	err = os.Remove("newdir")
	if err != nil {
		log.Fatal(err)
	}
}
