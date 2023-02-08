package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//设置文件上传大小限制,默认是32MB
	router.MaxMultipartMemory = 64 << 20 // 64MB
	router.POST("/upload", func(ctx *gin.Context) {
		//file 是表单字段名字
		file, _ := ctx.FormFile("file")
		//打印上传的文件名
		log.Println(file.Filename)

		//将上传的文件保存到./data/dusha/ganduward.jpg
		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!"), file.Filename)
	})
	router.Run(":8088")
}
