package main

import "fmt"

func randGenerator() chan int {
	ch := make(chan int)

	go func() {
		for {
			//select会尝试执行各个case如果都可以执行就随机选择一个执行
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	return ch
}

func main() {
	generator := randGenerator()

	for i := 0; i < 10; i++ {
		fmt.Println(<-generator)
	}
}
