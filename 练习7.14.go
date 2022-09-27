/*
方法一
package main

import "fmt"

	func main() {
		var str = []byte("Google")
		var a []byte
		for i := len(str) - 1; i >= 0; i-- {
			a = append(a, str[i])
		}
		fmt.Println(string(a))
	}
*/
//方法2
package main

import "fmt"

func main() {
	var str = []byte("Google")
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
	}
	fmt.Println(string(str))
}
