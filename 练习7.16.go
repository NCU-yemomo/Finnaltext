package main

import "fmt"

func main() {
	var a []int = []int{3, 7, 8, 6, 1, 8, 92, 4, 3, 8, 1}
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
