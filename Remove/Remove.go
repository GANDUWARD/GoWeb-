package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("dir1", 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Remove("dir1")
	if err != nil {
		log.Fatal(err)
	}
	//创建多级目录
	os.MkdirAll("test1/test2/test3", 0777)
	//删除test1目录及其子目录
	err = os.RemoveAll("test1")
	if err != nil {
		log.Fatal(err)
	}
}
