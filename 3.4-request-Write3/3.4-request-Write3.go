package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	//返回JSON格式数据
	greeting := Greeting{
		"欢迎一起学习Go Web编程",
	}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func main() {
	http.HandleFunc("/hello", Hello)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
