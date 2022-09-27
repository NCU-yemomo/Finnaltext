package main

import "fmt"

func main() {
	type T struct {
		a int
	}
	var s T        //实例化对象s
	fmt.Print(s.a) //Go的类不能动态添加

}
