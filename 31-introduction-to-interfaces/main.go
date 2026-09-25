package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func printArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	r := Rectangle{width: 10, height: 5}
	c := Circle{radius: 5}

	printArea(r) // 50
	printArea(c) // 78.53981633974483
}
