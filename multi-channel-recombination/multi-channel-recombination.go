package main

import (
	"fmt"
	"math/rand"
	"time"
)

//这个函数可以用来处理比较耗时的事情例如计算

func doCompute(x int) int {
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Millisecond))) //模拟计算
	return 1 + x                                                     //假设该计算为很费时的计算
}

// 每个分支开出一个goroutine来做计算,并把计算结果发送到各自通道中
func branch(x int) chan int {
	ch := make(chan int)
	go func() {
		ch <- doCompute(x)
	}()
	return ch
}

func Recombination(chs ...chan int) chan int {
	ch := make(chan int)
	/*for _, c := range chs {
		//注意此处要明确传值
		go func(c chan int) {
			ch <- <-c
		}(c) //复合，就是ch通道先接受通道c发送的数据，再自己发送
	}
	*/
	//select语句
	go func() {
		for i := 0; i < len(chs); i++ {
			select {
			case v := <-chs[i]:
				ch <- v
			}
		}
	}()

	return ch
}

func main() {
	//返回复合后的结果
	result := Recombination(branch(10), branch(20), branch(30))

	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
