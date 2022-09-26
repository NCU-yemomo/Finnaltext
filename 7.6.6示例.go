package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []int{3, 5, 7, 2, 3, 4, 9, 1, 0, 3, 88, 0, 1, 10, 2, 45}
	sort.Ints(a)
	fmt.Print(a)
}
