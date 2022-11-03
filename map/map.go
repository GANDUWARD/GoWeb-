package main

import "fmt"

func main() {
	achievement := map[string]float32{
		"zhangsan": 99.5, "xiaoli": 88,
		"wangwu": 96, "lidong": 100,
	}
	fmt.Println(achievement)
	map1 := make(map[int][]int)
	fmt.Println(map1)
}
