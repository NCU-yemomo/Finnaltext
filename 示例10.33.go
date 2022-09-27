package main

import "fmt"

type number struct {
	f float32
	a int
}

type nr number

func main() {
	a := number{5.0, 1}
	b := nr{5.0, 1}
	var c = number(b)
	fmt.Print(a, b, c)
}
