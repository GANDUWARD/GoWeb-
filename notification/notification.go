package main

import "fmt"

func SendNotification(user string) chan string {
	//此处省略查询数据库获取新消息
	//声明一个通道来保存信息
	notifications := make(chan string, 500)

	//开启一个通道
	go func() {
		//将消息放入通道
		notifications <- fmt.Sprintf("Hi %s, welcome to our site!", user)
	}()

	return notifications
}

func main() {
	barry := SendNotification("barry")         //获取barry的消息
	ganduward := SendNotification("ganduward") //获取ganduward的消息

	//将获取的消息返回
	fmt.Println(<-barry)
	fmt.Println(<-ganduward)
}
