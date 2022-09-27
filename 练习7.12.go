package main

import "fmt"

func main() {
	a, b := myfunc("hello", 2)
	fmt.Println(a)
	fmt.Println(b)

}
func myfunc(str string, i int) (str1 string, str2 string) {
	str1 = str[:i+1]
	str2 = str[i+1:]
	return
}
