package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("123")
	buffer.WriteString("hello")
	a := buffer.String()
	fmt.Print(a)
}
