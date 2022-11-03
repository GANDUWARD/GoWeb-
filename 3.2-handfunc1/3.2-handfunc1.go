package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome GANDUWARD!")
}

func main() {
	http.HandleFunc("/", hi)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
