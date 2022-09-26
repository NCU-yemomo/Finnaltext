package main

import "fmt"

func main() {
	s := "hello"
	var a = []byte(s)
	a[0] = 'c'
	s_new := string(a)
	fmt.Printf(s_new)
}
