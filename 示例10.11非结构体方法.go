package main

import "fmt"

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	var a = IntVector{2, 5, 6, 8, 4, 9}
	fmt.Println(a.Sum())
	//类似于python自带的列表sum方法
}
