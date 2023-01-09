package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

// 参数结构体
type ArgsTwo struct {
	X, Y int
}

/*DialHTTP有问题*/
func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8088")
	if err != nil {
		log.Fatal("在这里发生错误了：DialHTTP", err)
	}
	//获取第一个输入值
	i1, _ := strconv.Atoi(os.Args[1])
	//获取第二个输入值
	i2, _ := strconv.Atoi(os.Args[2])
	args := ArgsTwo{i1, i2}
	var reply int
	//调用命名函数，等待它完成，并返回其错误状态
	err = client.Call("Algorithm.Sum", args, &reply)
	if err != nil {
		log.Fatal("Call Sum algorithm error:", err)
	}
	fmt.Printf("Algorithm和为:%d+%d=%d\n", args.X, args.Y, reply)
}
