package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("open.TXT")
	if err != nil {
		fmt.Printf("打开文件出错:%v\n", err)
	}
	fmt.Println(file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Printf("关闭文件出错：%v\n", err)
	}
}
