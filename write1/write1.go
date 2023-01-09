package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("write1.TXT", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content := []byte("阿斯顿法国红酒看来")
	if _, err := file.Write(content); err != nil {
		fmt.Println(err)
	}
	fmt.Println("写入成功！")
}
