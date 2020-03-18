package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
}

func main() {
	triangle := triangle{base: 5, height: 3}
	printArea(triangle)

	square := square{side: 3}
	printArea(square)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}
