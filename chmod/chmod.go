package main

import (
	"fmt"
	"os"
)

func main() {
	//Create()函数会根据传入的文件名创建文件，默认是0666
	fp, err := os.Create("./chmod1.TXT")
	if err != nil {
		fmt.Println("文件创建失败！")
	}
	defer fp.Close()
	fileinfo, err := os.Stat("./chmod1.TXT")
	if err != nil {
		fmt.Println(err)
	}
	filemode := fileinfo.Mode()
	fmt.Println(filemode)
	os.Chmod("./chmod1.TXT", 0777) //通过chmod函数重新赋权限-rwxrwxrwx
	fileinfo, err = os.Stat("./chmod1.TXT")
	if err != nil {
		fmt.Println(err)
	}
	filemode = fileinfo.Mode()
	fmt.Println(filemode)
}
