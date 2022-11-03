package main

import (
	"fmt"
	"net/http"
)

func Redict(w http.ResponseWriter, r *http.Request) {
	//设置一个301重定向
	w.Header().Set("Location", "https://www.4399.com")
	w.WriteHeader(301)
}

func main() {
	http.HandleFunc("/redirect", Redict)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
