package main

import "fmt"

func main() {
	var a []int = []int{8, 6, 9, 8, 2, 4, 3, 5, 7, 1, 2, 5, 0, 4, 4, 7, 8, 89, 2}
	fmt.Println(minSlice(a))
}
func minSlice(a []int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}
