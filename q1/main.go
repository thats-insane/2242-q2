package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	bigTriangle := Triangle{
		Base:   3,
		Height: 4,
	}
	smallCircle := Circle{
		Radius: 10,
	}

	printArea(bigTriangle)
	printArea(smallCircle)
}
