package main

import "fmt"

type employee struct {
	salary int
}

func main() {
	var percent float32
	var salary int
	fmt.Println("请输入您想增加的薪水百分比（单位%")
	fmt.Scan(&percent)
	fmt.Println("请输入您的薪水")
	fmt.Scan(&salary)
	var employee1 = employee{salary}
	fmt.Println(employee1.giveRaise(percent))
}
func (self *employee) giveRaise(a float32) float32 {
	return float32((self.salary)) * (a/100 + 1)
}
