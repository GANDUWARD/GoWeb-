package main

import (
	"fmt"
	"os"
)

func main() {
	//以读写方式打开文件
	fp, err := os.OpenFile("./open.TXT", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败！")
		return
	}

	//defer延迟调用
	defer fp.Close()
}
