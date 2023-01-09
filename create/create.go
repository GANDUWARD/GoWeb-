package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	//Create()函数会根据传入的文件名创建文件，默认权限是0666
	fp, err := os.Create("./demo.TXT") //如果文件已经存在，则将文件清空
	fmt.Println(fp, err)
	fmt.Printf("%T", fp) //*os.File 文件的指针类型

	if err != nil {
		fmt.Println("文件创建失败！")
		//创建失败的原因有：
		//1、路径不存在 2、权限不足 3、打开文件数量超出上限 4、磁盘空间不足等
		return
	}
	//defer 延迟调用
	defer fp.Close() //关闭文件，释放资源
}
