package main

import "fmt"

type inners struct {
	in1 int
	in2 int
}

type outers struct {
	in3 int
	f1  float32
	string
	inners
}

func main() {
	var a outers
	a.in2 = 3
	a.in1 = 4
	a.in3 = 5
	a.f1 = 3.14
	a.string = "hello"
	fmt.Println(a)
	fmt.Println(a.in1)
	fmt.Println(a.in2)

	outer2 := outers{1, 3.14, "hello", inners{2, 3}}
	fmt.Println(outer2)
}
