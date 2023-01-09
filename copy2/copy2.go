package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 自定义复制函数
func DoCopy(srcFileName string, dstFileName string) {
	//打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		log.Fatalf("源文件读取失败,err:%v\n", err)
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			log.Fatalf("源文件关闭失败，err:%v\n", err)
		}
	}()

	//创建目标文件,稍后会向这个目标文件写入复制的内容
	distFile, err := os.Create(dstFileName)
	if err != nil {
		log.Fatalf("目标文件创建失败,err:%v\n", err)
	}
	defer func() {
		err = distFile.Close()
		if err != nil {
			log.Fatalf("目标文件关闭失败，err:%v\n", err)
		}
	}()
	//定义指定长度的字节切片，每次最多读取指定长度
	var tmp = make([]byte, 1024*4)
	for {
		n, err := srcFile.Read(tmp)
		n, _ = distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF {
				return
			} else {
				log.Fatalf("复制过程中发生错误,错误err:%v\n", err)
			}
		}
	}
}
func main() {
	//创建一个test.zip文件
	src := "./test.zip"
	_, err := os.Create(src)
	if err != nil {
		fmt.Println(err)
	}
	//复制一个名为test2.zip的文件
	dst := "./test2.zip"
	DoCopy(src, dst)
}
