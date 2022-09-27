package main

import "fmt"

type hello struct {
	thing int
}

func (b *hello) muti10() {
	b.thing = b.thing * 10
}
func (b hello) print() int {
	return b.thing
}
func main() {
	var b = hello{10}
	fmt.Println(b.print())
	b.muti10()
	fmt.Println(b)
	fmt.Println(b.print())
}
