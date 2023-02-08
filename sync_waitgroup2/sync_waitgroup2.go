package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testFunc := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("%v goroutine start ...\n", id)
		time.Sleep(2)
		fmt.Printf("%v goroutine exit ...\n", id)
	}
	var wg sync.WaitGroup
	const N = 3
	wg.Add(N)
	for i := 0; i < 3; i++ {
		go testFunc(&wg, i)
	}
	fmt.Println("Waiting for all goroutines")
	wg.Wait()
	fmt.Println("All goroutine finished!")
}
