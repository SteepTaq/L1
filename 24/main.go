package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	x1 := p.x - other.x
	y1 := p.y - other.y
	return math.Hypot(x1, y1)
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	p3 := NewPoint(-6, 4)

	fmt.Printf("точка 1: %v\n", p1)
	fmt.Printf("точка 2: %v\n", p2)
	fmt.Printf("точка 3: %v\n", p3)
	fmt.Println()
	fmt.Printf("расстояние от %v до %v: %v\n", p1, p2, p1.Distance(p2))
	fmt.Printf("расстояние от %v до %v: %v\n", p1, p3, p1.Distance(p3))
	fmt.Printf("расстояние от %v до %v: %v\n", p2, p3, p2.Distance(p3))
}
