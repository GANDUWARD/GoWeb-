package main

import "fmt"

type server struct {
	number_of_data int
}

func (x *server) SetNumber_of_int(num int) {
	x.number_of_data = num
}

type user struct {
	Situation string
}
type plot struct {
	r int
	g int
	b int
}

func (x *plot) init(red int, green int, blue int) {
	x.r = red
	x.g = green
	x.b = blue
}
func (x *plot) Set_r(num int) {
	x.init(num, 0, 0)
}
func (x *plot) Set_g(num int) {
	x.init(0, num, 0)
}
func (x *plot) Set_b(num int) {
	x.init(0, 0, num)
}

type set interface {
	Set_r(num int)
	Set_g(num int)
	Set_b(num int)
}
type dusha struct {
	Behind server
	Unit   user
	Front  plot
	method set
}

func main() {
	var SyS dusha
	fmt.Println(SyS)
	SyS.Behind.SetNumber_of_int(10)
	SyS.Front.init(230, 230, 230)
	SyS.Unit.Situation = "dusha"
	fmt.Println(SyS)

}
