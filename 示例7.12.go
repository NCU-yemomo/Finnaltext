package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Println(n, slTo)

	SlChange := []int{1, 3, 4}
	a := copy(SlChange, slFrom) //覆盖多少数就是多少数，是完全覆盖
	fmt.Println(a, SlChange)

	SlChange2 := []int{1, 2}
	b := copy(SlChange2, slFrom)
	fmt.Println(b, SlChange2)

	sl3 := append(SlChange2, SlChange...) //...的作用其实就是扩宽
	fmt.Println(sl3)

}
