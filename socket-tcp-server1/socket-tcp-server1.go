package main

import (
	"fmt"
	_ "io"
	"log"
	"net"
)

func Server() {
	//用Listen()函数创建服务器端
	//tcp : 网络协议
	//本机的ip地址及端口号:127.0.0.1:8088
	l, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close() //延时关闭
	//循环等待客户端访问
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("访 问 客 户 端 信 息 : con=%v 客 户 端 ip=%v\n", conn, conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

// 服务器端处理从客户端接收的数据
func handleConnection(c net.Conn) {
	defer c.Close() //关闭conn

	for {
		//1.等待客户端通过conn对象发送信息
		//2.如果客户端没有发送数据，则goroutine就堵塞在这里
		fmt.Printf("服务器在等待客户端%s发送信息\n", c.RemoteAddr().String())
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		//3.显示客户端发送到服务器端的内容
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	Server()
}
