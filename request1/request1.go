package main

import "net/url"

func main() {
	path := "http://localhost:8082/article?id=1"
	p, _ := url.Parse(path) //解析URL
	println(p.Host)
	println(p.User)
	println(p.RawQuery)
	println(p.RequestURI())
}
