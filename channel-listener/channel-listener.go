package main

import "fmt"

func foo(i int) chan int {
	ch := make(chan int)
	go func() {
		ch <- i
	}()
	return ch
}

func main() {
	ch1, ch2, ch3 := foo(3), foo(6), foo(9)

	ch := make(chan int)

	//开启一个goroutine监视各个通道数据输出,并收集数据到通道ch中
	go func() {
		for {
			//监视通道ch1 ch2 ch3的输出,并将其全部输入通道ch中
			select {
			case v1 := <-ch1:
				ch <- v1
			case v2 := <-ch2:
				ch <- v2
			case v3 := <-ch3:
				ch <- v3
			}
		}
	}()

	//阻塞主线,取出通道ch中的数据
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
}
