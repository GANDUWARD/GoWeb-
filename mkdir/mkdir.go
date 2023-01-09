package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//创建一个名为test目录，权限为0777
	err := os.Mkdir("test", 0777)
	if err != nil {
		fmt.Println(err)
	}
	//创建一个多级子目录，例如dir1/dir2/dir3
	err1 := os.MkdirAll("dir1/dir2/dir3", 0777)
	if err1 != nil {
		fmt.Println(err1)
	}
	uploadDir := "static/upload" + time.Now().Format("2006/01/02/")
	err2 := os.MkdirAll(uploadDir, 0777)
	if err2 != nil {
		fmt.Println(err2)
	}
}
