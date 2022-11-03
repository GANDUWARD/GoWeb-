package main

import (
	"fmt"
	"log"
	"net/http"
)

type WelcomHandler struct {
	Language string
}

// 定义一个ServeHTTP方法，以实现Handler接口
func (h WelcomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", h.Language)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cn", WelcomHandler{Language: "欢迎一起来学Go Web!"})
	mux.Handle("/en", WelcomHandler{Language: "Welcome you,let's learn Go Web!"})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
