package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
} //匿名类型
//嵌套来定义方法

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}
func main() {
	m := myTime{time.Now()}
	//开始调用匿名Time的的String方法
	fmt.Println(m.String())
	//
	fmt.Println(m.first3Chars())
}
