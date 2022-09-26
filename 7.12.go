package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Remove(s, 1, 6))
}
func Remove(s []int, start int, end int) []int {
	s = append(s[:start], s[end+1:]...)
	return s
}
