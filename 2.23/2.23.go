package main

import (
	"net/http"
)

type Refer struct {
	handler http.Handler
	refer   string
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}
func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(403)
	}
}
func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shirdon.com",
	}
	http.HandleFunc("/Hello", SayHello)
	http.ListenAndServe(":8080", referer)
}
