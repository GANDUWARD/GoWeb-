package main 

import(
	"fmt"
)

type User struct{
	Uid int
	Name string
	Phone string
	Email string
	Password string
}

//查询多条数据
func queryMultiRow()([]User){
	rows, err := db.
}