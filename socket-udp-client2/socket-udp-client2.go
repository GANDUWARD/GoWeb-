package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8023")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	//通过go关键字启动goroutine，从而支持并发
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) //从键盘读取内容放到str字符串里
			if err != nil {
				fmt.Println("os.Stdin.err=", err)
				return
			}
			conn.Write(str[:n]) //发送给服务器
		}
	}()
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("服务器发来:", string(buf[:n]))
	}
}
