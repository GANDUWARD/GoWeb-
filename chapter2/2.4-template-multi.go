package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 定义一个UserInfo结构体
type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func tmplSample(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./t.html", "./ul.html")
	if err != nil {
		fmt.Println("create template failed,err:", err)
		return
	}
	user := UserInfo{
		Name:   "张三",
		Gender: "男",
		Age:    28,
	}
	tmpl.Execute(w, user)
}

func main() {
	http.HandleFunc("/", tmplSample)
	http.ListenAndServe(":8087", nil)
}
