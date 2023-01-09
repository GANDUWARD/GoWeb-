package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.OpenFile("read.TXT", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("打开文件出错：%v\n", err)
	}
	//及时关闭文件句柄
	defer file.Close()
	//bufio.NewReader(rd io.Reader) *Reader
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}
