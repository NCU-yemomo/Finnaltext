package main

import "fmt"

type Log struct {
	msg string
}
type Customer struct {
	Name string
	Log  //这叫继承，前面是组合，这个其实更方便
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
func (l *Log) String() string {
	return l.msg
}
func main() {
	var a Customer
	a.Name = "yemmomo"
	a.msg = "1-今天我要吃饭"
	a.Add("2-今天我还要写作业")
	fmt.Print(a.String())
}
