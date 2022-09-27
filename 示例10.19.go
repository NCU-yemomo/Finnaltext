package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Baral Obama"
	c.log = new(Log) //指针联系，这叫聚合！
	c.log.msg = "1-Yes we can"
	//shorter
	c = &Customer{"Baral Obama", &Log{"1-Yes we can"}}
	//日志添加
	c.log.Add("2-After me the world willl be a better place")
	//打印日志
	fmt.Print(c.log.String())
}
func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}
