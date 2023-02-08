package main

import (
	"fmt"
)

func main() {

	ch, quit := make(chan int), make(chan int)

	go func() {
		ch <- 8
		quit <- -1
	}()
	for isQuit := false; !isQuit; {
		select {
		case v := <-ch:
			fmt.Printf("received %d from ch", v)
		case <-quit:
			isQuit = true
		}
	}
}
