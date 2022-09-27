package main

import "fmt"

type T struct {
	a int
	b float32
}

func main() {
	var s = T{a: 3, b: 3.4}
	fmt.Print(s.a, s.b)
}
