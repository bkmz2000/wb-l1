package main

import "math"

type Point struct {
	x, y float64
}

func (p Point) Dist(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y

	return math.Sqrt(dx*dx + dy*dy)
}
