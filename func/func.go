package main

import "fmt"

func dusha(array []int) (m int) {
	m = array[0]
	for _, v := range array {
		if v < m {
			m = v
		}
	}
	return
}
func gougu(x, y int) int {
	return (x*x + y*y)
}
func sum(arg ...int) (sum int) {
	for _, value := range arg {
		sum += value
	}
	return
}
func main() {
	var x = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var min int
	min = dusha(x)
	fmt.Printf("最小值是:%d\n", min)
	a := 3
	b := 4
	var c int
	c = gougu(a, b)
	fmt.Println("c为", c)
	var s = 0
	s = sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(s)
}
