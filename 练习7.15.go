package main

import "fmt"

func main() {
	var a = []byte("hellookkjp")
	var str []byte
	for i, j := range a {
		if i != len(a)-1 {
			if j == a[i+1] {
				continue
			} else {
				str = append(str, j)
			}
		} else {
			str = append(str, j)
		}
	}
	fmt.Println(string(str))
}
