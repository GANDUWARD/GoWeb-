package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	conn.Send("HSET", "students", "name", "jim", "age", "19")
	conn.Send("HSET", "students", "scores", "100")
	conn.Send("HGET", "students", "age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v \n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%v \n", res3)
}
