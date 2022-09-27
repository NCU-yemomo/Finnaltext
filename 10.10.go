package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}
func (b *Base) SetId(n int) {
	b.id = n
}

type Person struct {
	FirstName string
	Base
	LastName string
}
type Employee struct {
	Person
	salary float64
}

func (e *Employee) RasingSalary(percent float64) {
	e.salary += e.salary * percent / 100
}
func main() {
	var a = new(Employee)
	var name string
	var salary float64
	var id int
	var percent float64
	fmt.Println("请输入您的名字")
	fmt.Scan(&name)
	a.FirstName = string(name[0])
	a.LastName = string(name[1:])
	fmt.Println("请输入您的id")
	fmt.Scan(&id)
	a.SetId(id)
	fmt.Println("请输入您的薪水")
	fmt.Scan(&salary)
	a.salary = salary
	fmt.Println("请输入您想长的倍率")
	fmt.Scan(&percent)
	a.RasingSalary(percent)
	fmt.Printf("姓名：%s %s id:%d 薪水：%.2f", a.FirstName, a.LastName, a.Id(), a.salary)
}
