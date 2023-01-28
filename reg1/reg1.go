package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I_Love_GANDUWARD"
	res := strings.Split(s, "_")
	for value := range res {
		fmt.Println(res[value])
	}
	value := "a|b|c|d"
	//分割成三部分
	result := strings.SplitN(value, "|", 3)
	for v := range result {
		fmt.Println(v)
	}
}
