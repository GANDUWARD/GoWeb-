package main

import (
	"fmt"
	"time"
)

func Timer(duration time.Duration) chan bool {
	ch := make(chan bool)

	go func() {
		time.Sleep(duration)
		//时间到就会发送真
		ch <- true
	}()
	return ch
}

func main() {
	//定时5s
	timeout := Timer(5 * time.Second)

	for {
		select {
		case <-timeout:
			fmt.Println("time up")
			return
		}
	}
}
