package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("writeAt.TXT")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("阿斯顿法国红酒看来")
	n, err := file.WriteAt([]byte("dusha！"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
