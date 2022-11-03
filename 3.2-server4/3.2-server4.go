package main

import (
	"log"
	"net/http"

	"gitee.com/shirdon/goWebActualCombat/chapter3/controller"
)

// gittee权限有误
func main() {
	http.HandleFunc("/getUser", controller.UserController{}.getUser)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
