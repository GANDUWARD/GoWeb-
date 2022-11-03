package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,err:", err)
		return
	}
	defer c.Close()

	c.Send("SET", "username1", "jim")
	c.Send("SET", "username2", "jack")
	c.Flush()

	v, err := c.Receive()
	fmt.Printf("v:%v,err:%v\n", v, err)
	v, err = c.Receive()
	fmt.Printf("v:%v,err:%v\n", v, err)
	v, err = c.Receive() //一直等待
	fmt.Printf("v:%v,err:%v\n", v, err)
}
