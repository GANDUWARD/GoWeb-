package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("你好~欢迎一起学习Go Web编程"))
}

func main() {
	http.HandleFunc("/welcome", Welcome)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
