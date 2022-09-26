package main

import "fmt"

func main() {
	var arrF []float32 = []float32{1.3, 5.6, 8.2, 4.6, 8.8}
	fmt.Println(Add(arrF...))
}

func Add(a ...float32) float32 {
	var result float32 = 0
	for _, i := range a {
		result += i
	}
	return result
}
