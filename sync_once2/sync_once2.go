package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func func1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func func2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- 2 * x
	}
	once.Do(func() { close(ch2) }) //确保某个操作只被执行一次
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	wg.Add(3)
	go func1(ch1)
	go func2(ch1, ch2)
	go func2(ch1, ch2)
	wg.Wait()

	for ret := range ch2 {
		fmt.Println(ret)
	}
}
