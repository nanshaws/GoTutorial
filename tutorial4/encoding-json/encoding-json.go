package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// 创建一个 User 对象
	user := User{ID: 1, Name: "Bob"}

	// 将 User 对象编码为 JSON 字符串
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData)) // 输出: {"id":1,"name":"Bob"}

	// 将 JSON 字符串解码回 User 对象
	var decodedUser User
	if err := json.Unmarshal(jsonData, &decodedUser); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded User: %+v\n", decodedUser) // 输出: Decoded User: {ID:1 Name:Bob}
}
