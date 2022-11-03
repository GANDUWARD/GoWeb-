package main

import (
	"fmt"
)

func visitPrint(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}
func dusha() {
	fmt.Println("defer语句现在执行!")
}
func main() {
	sli := []int{1, 6, 8}
	defer dusha()
	//使用匿名函数打印切片内容
	visitPrint(sli, func(value int) {
		fmt.Println(value)
	})
}
