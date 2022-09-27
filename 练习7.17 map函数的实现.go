package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(map_(mulity, a))
}

func map_(F func(n int) int, nums []int) []int {
	for i, j := range nums {
		nums[i] = F(j)
	}
	return nums
}
func mulity(n int) int {
	return n * 10
}
