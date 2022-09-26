package main

import "fmt"

func main() {
	var s []int = []int{1, 2, 3, 4, 5}
	var factor int = 2
	result := len(s) * factor
	s = append(s, make([]int, result)...)
	fmt.Println(s)
}
