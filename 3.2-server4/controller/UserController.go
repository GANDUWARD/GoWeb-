package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"gitee.com/shirdon/goWebActualCombat/chapter3/model"
)

type UserController struct {
}

func (c UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	uid, _ := strconv.Atoi(query["uid"][0])

	//此处调用模型从数据库中获取数据
	user := model.GetUser(uid)
	fmt.Println(user)

	t, _ := template.ParseFiles("view/t3.html")
	userInfo := []string{user.Name, user.Phone}
	t.Execute(w, userInfo)
}
