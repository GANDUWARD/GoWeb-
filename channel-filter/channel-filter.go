package main

import (
	"fmt"
)

// 生成自增的整数
func IntegerGenerator() chan int {
	var ch chan int = make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i //直到通道索要数据才把i添加进通道
		}
	}()
	return ch
}

func Filter(in chan int, number int) chan int {
	//输入一个整数队列,筛出的是number的倍数
	//将不是number的倍数放入输入队列中
	out := make(chan int)

	go func() {
		for {
			i := <-in //从输入中取一个数
			if i%number != 0 {
				out <- i //将数放入通道中
			}
		}
	}()
	return out
}

func main() {
	const max = 100
	numbers := IntegerGenerator()
	number := <-numbers //从生成器抓取整数2

	for number <= max {
		fmt.Println(number)               //打印素数
		numbers = Filter(numbers, number) //筛出number的倍数
		number = <-numbers
	}
}
