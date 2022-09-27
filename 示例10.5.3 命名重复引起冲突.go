package main

type A struct {
	a int
}
type B struct {
	a, b int
}
type C struct {
	A
	B //起了冲突
}

func main() {
	var c C
	c.a = 3 //error:不明确的引用a

}
