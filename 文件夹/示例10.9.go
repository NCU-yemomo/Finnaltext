package main

import "fmt"

type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}

func main() {
	n := B{A{1, 2}, 3.4, 4.4}
	fmt.Println(n.ax, n.ay, n.by, n.bx)
	fmt.Println(n.A)
}
