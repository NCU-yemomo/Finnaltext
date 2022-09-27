package main

import "fmt"

//看看append方法在go如何实现

type List []int

func (self *List) append(factor int) []int {
	*self = append(*self, factor)
	return *self
}
func (l List) len() int {
	return len(l)
}
func main() {
	var a List
	for i := 0; i < 11; i++ {
		a.append(i)
	}
	fmt.Println(a)
	fmt.Println(a.len())
	//定义在指针和原先玩意的方法都可以直接被调用，调用时会发生转换
	//很方便吧~一般指针就要对数组做出修改了，推荐使用指针，节省内存
}
