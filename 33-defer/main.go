package main

import "fmt"

func process() {
	fmt.Println("First output")
	defer fmt.Println("Cleaning up")
	fmt.Println("Second output")
}

func demo() {
	x := 10
	defer fmt.Println("x:", x)
	x = 20
	fmt.Println("x:", x)
}

func show() {
	for i := range 3 {
		defer fmt.Println("i:", i)
	}
}

func main() {
	process()
	demo()
	show()
}
