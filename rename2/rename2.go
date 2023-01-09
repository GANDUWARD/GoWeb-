package main

import (
	"fmt"
	"os"
)

func main() {
	//创建一个名为"test_rename.TXT"的空文件
	f, err := os.Create("./test_rename.TXT") //如果文件已存在，则将文件清空
	if err != nil {
		fmt.Println(err)
	}
	//创建一个名为test_rename的目录，权限为0777
	f.Close()
	err = os.Mkdir("test_rename", 0777)
	//注意要将文件关闭后才能被移动
	//将test_rename.TXT移动到test_rename目录下，并将其重命名为test_rename_new.TXT
	err = os.Rename("./test_rename.TXT", "./test_rename/test_rename_new.TXT")
	if err != nil {
		fmt.Println(err)
		return
	}
}
