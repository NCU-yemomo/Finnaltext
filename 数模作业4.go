package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums :=[]int
	for i := 0; i < n+1; i++ {


	}
	result := 0
	for _, i := range nums {
		result += i
	}
	if result%2 == 0 {
		fmt.Printf("Bob")
	} else {
		fmt.Printf("Alice")
	}
}
