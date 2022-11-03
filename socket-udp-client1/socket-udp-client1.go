package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("你好，我是用UDP的客户端"))

	buf := make([]byte, 1024)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		return
	}
	fmt.Println("服务器发来：", string(buf[:n]))
}
