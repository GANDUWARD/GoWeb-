package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为"test_rename"的目录，权限为0777
	err := os.Mkdir("test_remove", 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("created dir:test_remove")
	//创建一个名为"test_remove1.TXT"的空文件
	f1, err1 := os.Create("./test_remove/test_remove1.TXT") //如果文件已经存在,将其清空
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("created dir:test_remove1.TXT")
	f2, err2 := os.Create("./test_remove/test_remove2.TXT") //如果文件已经存在,将其清空
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("created dir:test_remove2.TXT")
	f3, err3 := os.Create("./test_remove/test_remove3.TXT") //如果文件已经存在,将其清空
	if err3 != nil {
		fmt.Println(err)
	}
	f1.Close()
	f2.Close()
	f3.Close()
	fmt.Println("created dir:test_remove3.TXT")
	err = os.Remove("./test_remove/test_remove1.TXT")
	if err != nil {
		fmt.Printf("removed ./test_remove/test_remove1.TXT err: %v \n", err)
	}
	fmt.Println("removed file:./test_remove/test_remove1.TXT")
	err = os.RemoveAll("./test_remove")
	if err != nil {
		fmt.Printf("remove all ./test_remove err: %v\n", err)
	}
	fmt.Println("removed all files:./test_remove")
}
