package main

import (
	"fmt"
	"github.com/Ki-rin/go-example/points"
)

func main() {
	objects := []points.Magnitude{
		points.Point{X: 3, Y: 4},  // Abs = 3*3 + 4*4 = 25
		points.Circle{Radius: 5},  // Abs = 5*5 = 25
		points.Point{X: 5, Y: 12}, // Abs = 25 + 144 = 169
	}

	for i, obj := range objects {
		fmt.Printf("Object #%d magnitude: %.2f\n", i+1, obj.Abs())
	}
}
