package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(p2 Point) float64 {
	r1 := math.Pow(p.X-p2.X, 2)
	r2 := math.Pow(p.Y-p2.Y, 2)

	return math.Sqrt(r1 + r2)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)

	fmt.Print(p1.Distance(p2))
}
