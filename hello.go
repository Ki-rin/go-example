package main

import (
	"fmt"
)

type Point struct {
	x, y float32
}

func Abs(point Point) float32 {
	return point.x*point.x + point.y*point.y
}

func (p Point) Abs() float32 {
	return p.x*p.x + p.y*p.y
}

func main() {

	p := Point{1, 2} // This is an assignment to the variable p
	fmt.Println(p.x)
	fmt.Println(p.Abs() != Abs(p))
}
