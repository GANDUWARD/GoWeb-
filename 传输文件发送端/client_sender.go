package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//	提示输入文件
	fmt.Println("请输入需要传输的文件：")
	var path string
	fmt.Scan(&path)
	//path = "info.jpg"
	//	获取文件名，info.Name
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}
	//主动连接服务器
	conn, err1 := net.Dial("tcp", "43.143.215.213:8088")
	if err1 != nil {
		fmt.Println("net.Dail err1:", err1)
		return
	}
	defer conn.Close()

	//给接收方先发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err", err2)
		return
	}

	//接受对方的回复，如果回复“ok”，说明对方准备好，可以发送文件，相当于一次简单的握手过程
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Write err:", err3)
	}
	if "ok" == string(buf[:n]) {
		//	发送文件内容
		SendFile(path, conn)
	}
}

func SendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	f, err4 := os.Open(path)
	if err4 != nil {
		fmt.Println("os.Open err:", err4)
		return
	}
	defer f.Close()
	//读文件内容，读多少内容就发送多少内容
	fmt.Println(path)
	buf := make([]byte, 1024*500)
	for {
		n, err5 := f.Read(buf) //从文件读取内容
		if err5 != nil {
			if err5 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err:", err5)
			}
			return
		}
		if n == 0 {
			fmt.Println("空文件发送！")
			return
		}
		//发送内容
		conn.Write(buf[:n]) //给服务器发送内容
	}
}
