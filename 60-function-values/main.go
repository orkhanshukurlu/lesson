package main

import "fmt"

func hello() string {
	return "Hello, World!"
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	f1 := hello()
	f2 := hello
	fmt.Println(f1)
	fmt.Println(f2())

	var a func(int, int) int
	a = add
	r := a(1, 2)
	fmt.Println(r)

	var b func(int, int) int
	b = add
	fmt.Println(b(5, 2))
	b = subtract
	fmt.Println(b(5, 2))
}
