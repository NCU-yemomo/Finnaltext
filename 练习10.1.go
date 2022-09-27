package main

import "fmt"

type Address struct {
	position string
	number   int
}

type VCard struct {
	name    string
	address *Address
	birth   [3]int
}

func main() {
	var _ = Address{"南昌", 214}
	var v VCard
	v.name = "陈卓"
	v.birth = [3]int{2004, 6, 12}
	fmt.Println(v.name, ",", v.address.position, ",", v.address.number, "-")
	for k, i := range v.birth {
		if k != 2 {
			fmt.Print(i, "-")
		} else {
			fmt.Print(i)
		}
	}

}
