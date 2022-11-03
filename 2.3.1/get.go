package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Print("err", err)
	}
	closer := resp.Body
	bytes, err := ioutil.ReadAll(closer)
	fmt.Println(string(bytes))
}
