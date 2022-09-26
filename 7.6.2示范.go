package main

import "fmt"

func main() {
	var a = []rune("hello")
	for _, i := range a {
		fmt.Printf("%c\n", i)
	}
	fmt.Println(len(a))
}
