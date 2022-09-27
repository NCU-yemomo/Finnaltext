package main

type Point struct {
	width  int
	length int
}

func (p *Point) Abs() int {
	return (p.width + p.length) * 2
}
func (p *Point) Scale() int {
	return p.width * p.length
}
