package main

import "fmt"

type Car struct {
	wheelCount int
}

func (self *Car) SetWheels(numbers int) {
	self.wheelCount = numbers
}

type Mercedes struct {
	Car
}

func (self *Mercedes) sayHiToMerkel() {
	fmt.Print("hello Merkel ", "you have ", self.wheelCount, " wheels")
}
func main() {
	var merkel = new(Mercedes)
	merkel.SetWheels(4)
	merkel.sayHiToMerkel()
}
