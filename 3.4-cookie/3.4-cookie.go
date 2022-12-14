package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("test_cookie")
	fmt.Printf("cookie:%#v,err:%v\n", c, err)

	cookie := &http.Cookie{
		Name:   "test_cookie",
		Value:  "krrsklHhefUUUFSSKLAkaLlJGGQEXZLJP",
		MaxAge: 3600,
		Domain: "localhost",
		Path:   "/",
	}
	http.SetCookie(w, cookie)

	//应在具体数据返回钱设置cookie,否则cookie设置不成功
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/", testHandle)
	err := http.ListenAndServe(":8086", nil)
	if err != nil {
		fmt.Println(err)
	}
}
