package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	html := `<html>
	<head>
	<title>用Weite()方法返回html文档</title>
	<body>
	<h1>你好，欢迎一起学习Go Web编程
	</body>
	<html>`
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
