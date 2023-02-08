package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 6; i <= 8; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()
	for receive := range ch {
		fmt.Println(receive)
		if receive == 8 {
			break
		}
	}
}
