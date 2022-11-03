package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func Welcome() string {
	return "Welcome"
}

func Doing(name string) string {
	// 有参数函数
	return name + ",learning Go Web template"
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	htmlByte, err := ioutil.ReadFile("./funcs.html")
	if err != nil {
		fmt.Println("read html failed,err:", err)
		return
	}
	// 自定义一个匿名模板函数
	loveGo := func() string {
		return "欢迎一起学习Go语言Web编程-------GANDUWARD"
	}
	//链式操作在Parse()方法之前调用Funcs()函数,用来添加自定义的loveGo函数
	tmp11, err := template.New("funcs").Funcs(template.FuncMap{"loveGo": loveGo}).Parse(string(htmlByte))
	if err != nil {
		fmt.Println("create template failed,err:", err)
		return
	}
	funcMap := template.FuncMap{
		//在FuncMap中声明要使用的函数,然后就能够在模板的字符串中使用该函数
		"Welcome": Welcome,
		"Doing":   Doing,
	}
	name := "GANDUWARD"
	tmp12, err := template.New("test").Funcs(funcMap).Parse("{{Welcome}}\n{{Doing .}}\n")
	if err != nil {
		panic(err)
	}

	//使用uesr渲染模板,并将结果写入w
	tmp11.Execute(w, name)
	tmp12.Execute(w, name)
}

/*func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe(":8087", nil)
}*/
