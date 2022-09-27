package main

import "fmt"

func main() {
	var num int
	var a []int
	var state bool
	var minute int
	var hour int
	for i := 0; i < 4; i++ {
		fmt.Scanf("%d", &num)
		a = append(a, num)
	}
	if a[3] < a[1] {
		minute = a[3] + 60 - a[1]
		state = false
	} else {
		minute = a[3] - a[1]
		state = true
	}
	if state {
		hour = a[2] - a[0]
	} else {
		hour = a[2] - a[0] - 1
	}
	fmt.Print(hour, minute)
}
