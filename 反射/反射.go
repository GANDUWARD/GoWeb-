package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	var (
		i     int = 3
		dusha int = 4
	)
	fmt.Println(i + dusha)
	fmt.Println("value:", reflect.ValueOf(i))
	var name interface{} = "shirdon"
	fmt.Printf("原始接口变量的类型为%T,值为%v \n", name, name)
	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)

	//从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为%T \n", t)
	fmt.Printf("从接口对象到反射对象：Value对象的类型为%T \n", v)
	//从反射对象到接口变量
	dudu := v.Interface()
	fmt.Printf("从反射对象到接口对象:新对象的类型为%T值为%v \n", dudu, dudu)
}
