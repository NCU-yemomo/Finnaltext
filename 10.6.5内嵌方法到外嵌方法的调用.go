package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func (p *point) abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y) //求模
}
func (p *point) SetValue(a float64, b float64) {
	p.x = a
	p.y = b
}

type NamePoint struct {
	point //继承
	name  string
}

func (self *NamePoint) abs(a int) float64 {
	return self.point.abs() * float64(a)
} //两个不同类型可以用同一个名字的方法
func main() {
	var a = new(NamePoint)
	a.SetValue(3, 4)
	a.name = "hello"
	fmt.Println(a.name, a.abs(100))
}
