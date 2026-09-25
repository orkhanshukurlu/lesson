package main

import "fmt"

func process() {
	fmt.Println("First output")
	defer fmt.Println("Cleaning up")
	fmt.Println("Second output")
}

func demo() {
	x := 10
	defer fmt.Println(x)
	x = 20
	fmt.Println(x)
}

func show() {
	for i := range 3 {
		defer fmt.Println(i)
	}
}

func main() {
	process()
	/*
		First output
		Second output
		Cleaning up
	*/
	demo()
	/*
		20
		10
	*/
	show()
	/*
		2
		1
		0
	*/
}
