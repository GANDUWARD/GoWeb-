package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 使用Go提供的net/rpc/jsonrpc
func init() {
	fmt.Println("JSON编码RPC,不是GOB编码,其他的和RPC概念一模一样")
}

type ArgsLanguage struct {
	Java, Go string
}

type Programmer string

func (m *Programmer) GetSkill(al *ArgsLanguage, skill *string) error {
	*skill = "Skill1:" + al.Java + ",Skill2" + al.Go
	return nil
}

func main() {
	//实例化
	str := new(Programmer)
	//注册服务
	rpc.Register(str)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8088")
	if err != nil {
		fmt.Println("ResolveTCPAddr err=", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("tcp listen err=", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
