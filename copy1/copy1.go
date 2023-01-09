package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//创建一个名为test_copy1.zip文件
	_, err := os.Create("./test_copy1.zip")
	if err != nil {
		fmt.Println(err)
	}
	//打开文件test_copy1.zip,获取文件指针
	srcFile, err := os.Open("./test_copy1.zip")
	if err != nil {
		fmt.Printf("open file err : %v\n", err)
		return
	}

	defer srcFile.Close()

	//打开文件要复制的新文件名test_copy2.zip,获取文件指针
	dstFile, err := os.OpenFile("./test_copy2.zip", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
	}
	defer dstFile.Close()

	//通过copy函数复制数据
	result, err := io.Copy(dstFile, srcFile)

	if err == nil {
		fmt.Println("复制成功,复制的字节数为：", result)
	}
}
