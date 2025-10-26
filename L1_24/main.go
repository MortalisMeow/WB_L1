package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) (float64, error) {
	if p == nil || other == nil {
		return 0, fmt.Errorf("nil point")
	}
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy), nil
}

func main() {
	point := NewPoint(-4.5, -2.5)
	otherPoint := NewPoint(-12.5, -2.5)

	dist, err := point.Distance(otherPoint)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(dist)
}
