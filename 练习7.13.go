package main

import "fmt"

func main() {
	var str string = "hello world"
	fmt.Println(str[len(str)/2:] + str[:len(str)/2])
}
