package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 定义一个初始化数据库函数
func initDB() (err error) {
	//链接数据库
	db, err = sql.Open("mysql", "root:a123456@tcp(127.0.0.1:3306)/ch4")
	if err != nil {
		return err
	}
	//尝试与数据库建立连接
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() //调用输出数据库函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
}
