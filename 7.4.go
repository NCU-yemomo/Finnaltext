package main

import "fmt"

func main() {
	fmt.Print(fib(10))
}
func fib(n int) []int {
	var a []int = make([]int, n)
	if n <= 1 {
		for i := range a {
			a[i] = 1
		}
		return a
	} else {
		a[0] = 1
		a[1] = 1
		for i := 2; i < len(a); i++ {
			a[i] = a[i-1] + a[i-2]
		}
		return a
	}
}
