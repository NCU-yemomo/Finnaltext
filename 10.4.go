package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func Area(p *Rectangle) int {
	return p.length * p.width
}
func Perimeter(p *Rectangle) int {
	return (p.length + p.width) * 2
}
func main() {
	var a = &Rectangle{3, 4}
	fmt.Print(Area(a))
	fmt.Print(Perimeter(a))

}
