package main

import (
	"fmt"
)

func main() {
	ch := make(chan string) //构建一个通道
	go func() {             //开启一个并发匿名函数
		fmt.Println("开始goroutine")
		ch <- "signal"
		fmt.Println("退出goroutine")
	}()
	fmt.Println("等待goroutine")
	<-ch //等待匿名goroutine
	fmt.Println("完成")
}
