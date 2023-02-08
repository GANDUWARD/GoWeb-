package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定义一个日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//可以通过上下文对象,设置一些依附在上下文对象里面的键/值数据
		c.Set("example", "hi!这是一个中间件数据")
		//在这里处理请求到达处理器函数之前的逻辑

		//调用下一个中间件，或者处理器的处理函数，具体得看注册了多少个中间件
		c.Next()

		//在这里可以处理返给客户端之前的响应逻辑
		latency := time.Since(t)
		log.Print(latency)

		//例如，查询请求状态码
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()
	//注册上面自定义的日志中间件
	r.Use(Logger())
	r.GET("/hi", func(c *gin.Context) {
		//在查询之前在日志中间件中注入的键值数据
		example := c.MustGet("example").(string)

		//打印
		log.Println(example)
	})

	//启动服务器端:0.0.0.0:8088
	r.Run(":8088")
}
