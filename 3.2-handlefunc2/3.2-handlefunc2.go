package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Web\n")
	fmt.Fprintf(w, "DUSHA!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hi)
	server := &http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
