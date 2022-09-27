package main

import "fmt"

type Camera struct {
}

func (c *Camera) TakeApicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func main() {
	cp := new(CameraPhone)
	fmt.Println(cp.TakeApicture())
	fmt.Println(cp.call())
}
