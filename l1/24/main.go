package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.y-p1.y, 2) + math.Pow(p2.x-p1.x, 2))
}

func main() {
	p1 := *NewPoint(1, 2)
	p2 := *NewPoint(4, 5)
	fmt.Println(Distance(p1, p2))
}
