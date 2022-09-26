package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var index = []int{10, 20, 30, 40}
	fmt.Println(Insert(s, index, 4))
}
func Insert(s []int, index []int, start int) []int {
	index = append(index, s[start+1:]...)
	s = append(s[:start+1], index...)
	return s
}
