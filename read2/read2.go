package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//用io/ioutil.ReadFile()函数一次性将文件读取到内存中
	filePath := "read2.TXT"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("读取文件出错:%v", err)
	}
	fmt.Printf("%v\n", content)
	fmt.Printf("%v\n", string(content))
}
