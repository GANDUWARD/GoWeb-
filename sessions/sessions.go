package main

import (
	//导入"github.com/gin-contrib/sessions"包
	"github.com/gin-contrib/sessions"
	//导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
	//导入gin框架包
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//创建基于cookie的存储引擎,password123456参数是用于加密的密钥
	store := cookie.NewStore([]byte("password123456"))
	//设置session中间件,参数my_session指的是session的名字,也是cookie的名字
	//store 是前面创建的存储引擎,可以将其替换成其他存储引擎
	r.Use(sessions.Sessions("my_session", store))
	r.GET("/hello", func(c *gin.Context) {
		//初始化session对象
		session := sessions.Default(c)

		//通过session.Get()函数读取session值
		//session是键值对格式数据,因此需要通过key查询数据
		if session.Get("hello") != "world" {
			//设置session数据
			session.Set("hello", "world")
			//删除session数据
			session.Delete("ganduward")
			//保存session数据
			session.Save()
			//删除整个session
			//session.Clear()
		}
		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})
	r.Run(":8088")
}

func Handler(c *gin.Context) {
	//根据cookie名字读取cookie值
	data, err := c.Cookie("my_cookie")
	if err != nil {
		//直接返回cookie值
		c.String(200, data)
		return
	}
	c.String(200, "not found!")
}
