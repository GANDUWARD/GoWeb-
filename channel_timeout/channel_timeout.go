package main

import (
	"fmt"
	"time"
)

func basic_channel(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- x
	}()
	return ch
}

func recombination() chan int {
	ch := make(chan int)
	ch1, ch2, ch3 := basic_channel(3), basic_channel(6), basic_channel(9)
	go func() {
		//timeout是一个计时通道，如果时间到会发信号
		timeout := time.After(1 * time.Second)
		for isTimeout := false; !isTimeout; {
			select {
			case v1 := <-ch1:
				fmt.Printf("received %d frome ch1\n", v1)
				ch <- v1
			case v2 := <-ch2:
				fmt.Printf("received %d frome ch1\n", v2)
				ch <- v2
			case v3 := <-ch3:
				fmt.Printf("received %d frome ch1\n", v3)
				ch <- v3
			case <-timeout:
				isTimeout = true
				<-timeout
			}
		}
	}()
	return ch
}
func main() {
	chs := recombination()
	for i := 0; i < 3; i++ {
		fmt.Println(<-chs)
	}
}
