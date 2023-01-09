package main

import (
	"fmt"
	"os"
)

func main() {
	//创建文件
	fp, err := os.Create("./link1.TXT")
	if err != nil {
		fmt.Println(err)
	}
	defer fp.Close()
	err = os.Link("link1.TXT", "link2.TXT")
	if err != nil {
		fmt.Println(err)
	}
}
