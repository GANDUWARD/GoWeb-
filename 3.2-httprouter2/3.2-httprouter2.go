package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Hostmap map[string]http.Handler

func (hs Hostmap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//先根据域名获取对应的Handler路由，然后调用处理(先发机制)
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func main() {
	userRouter := httprouter.New()
	userRouter.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("sub1"))
	})
	dataRouter := httprouter.New()
	dataRouter.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("sub2"))
	})
	//分别处理不同的域名
	hs := make(Hostmap)
	hs["sub1.localhost:8888"] = userRouter
	hs["sub2.localhost.8888"] = dataRouter

	log.Fatal(http.ListenAndServe(":8888", hs))
}
