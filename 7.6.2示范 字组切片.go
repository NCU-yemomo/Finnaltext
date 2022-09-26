package main

import "fmt"

func main() {
	a := "hello"[1:5]
	fmt.Println(a)
	//字符串可以遍历，不过出来的是unicode值
	for _, i := range a {
		fmt.Println(string(i))
	}
}

//因此，您必须先将字符串转换成字节数组，然后再通过修改数组中的元素值来达到修改字符串的目的，
//最后将字节数组转换回字符串格式。
