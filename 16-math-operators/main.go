package main

import (
	"fmt"
	"math"
)

func main() {
	a := 7
	b := 3
	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)

	fmt.Println("sqrt ", math.Sqrt(16))
	fmt.Println("pi ", math.Pi)
}
