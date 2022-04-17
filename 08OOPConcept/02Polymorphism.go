package main

import (
	"fmt"
	"math"
)

type shape01 interface {
	area() float64
}

type triangle struct {
	a float64
	h float64
}

type square struct {
	a float64
}

type rectangle01 struct {
	a float64
	b float64
}

func (t triangle) area() float64 {

	return (t.a * t.h) / 2
}

func (s square) area() float64 {

	return math.Pow(s.a, 2)
}

func (r rectangle01) area() float64 {

	return r.b * r.a
}

func printArea(shapes ...shape01) {
	for _, shape01 := range shapes {
		fmt.Println("Area: ", shape01.area())
	}
}

func main() {

	t := triangle{
		3, 4,
	}
	s := square{
		3,
	}
	r := rectangle01{
		3, 4,
	}

	printArea(t, s, r)
}
