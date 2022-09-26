package main

import "fmt"

func main() {
	var a int = 1
	var b float32 = 2.8
	fmt.Println(SumAndAverage(a, b))
}
func SumAndAverage(a int, b float32) (sum float32, average float32) {
	sum = float32(a) + b
	average = (float32(a) + b) / 2
	return
}
