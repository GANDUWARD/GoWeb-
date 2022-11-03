package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//访问静态文件
	router.ServeFiles("/static/*filepath", http.Dir("./files"))
	log.Fatal(http.ListenAndServe(":8086", router))
}
