// 可以多看看，是个很好的思想
package main

import "fmt"

func main() {
	var nums = []int{1, 4, 6, 8, 3, 9, 4, 6, 2, 3, 2, 8, 4}
	fmt.Println(Filter(nums, fn))
}
func Filter(s []int, fn func(int) bool) []int {
	var a = []int{}
	for _, i := range s {
		if fn(i) {
			a = append(a, i)
		}
	}
	return a
}
func fn(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}
