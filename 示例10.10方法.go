package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	//始终是定义一个类（结构体），结构体能够有多个实例化对象
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Println(two1.a, two1.b)
	fmt.Println(two1.AddThem())
	fmt.Println(two1.AddToParam(12))
}

func (self *TwoInts) AddThem() int {
	return self.a + self.b
}

// 类似于Python
// def add(self):
//
//	retrun self.a+self.b(这里self名称可以自己定义)
func (self *TwoInts) AddToParam(param int) int {
	return self.a + self.b + param
}
