package main

import (
	"fmt"
)

func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch) //管道可关闭
}

func main() {
	ch := make(chan int, 6)
	go fibonacci(cap(ch), ch) //cap()表示求通道的容量
	for j := range ch {
		fmt.Println(j)
	}
}
