package main

import "fmt"

func main() {
	var n int
	var sub1, sub2, sub3 int
	var state bool = false
	fmt.Scan(&n)
	for i := 10000; i <= 30000; i++ {
		sub1 = i / 100
		sub2 = (i % 10000 / 10)
		sub3 = i % 1000
		if sub1%n == 0 && sub2%n == 0 && sub3%n == 0 {
			fmt.Println(i)
			state = true
		}
	}
	if !state {
		fmt.Printf("No")
	}
}
