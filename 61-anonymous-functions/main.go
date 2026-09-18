package main

import "fmt"

func main() {
	msg := func(name string) {
		fmt.Println("Hello:", name)
	}

	msg("John")

	func(name string) {
		fmt.Println("Hello:", name)
	}("Alice")

	var m func(string)
	m = func(name string) {
		fmt.Println("Hello:", name)
	}
	m("Bob")

	s := []func(int, int) int{
		func(a, b int) int {
			return a + b
		},
		func(a, b int) int {
			return a - b
		},
		func(a, b int) int {
			return a * b
		},
		func(a, b int) int {
			return a / b
		},
	}

	fmt.Println(s[0](10, 2))
	fmt.Println(s[1](10, 2))
	fmt.Println(s[2](10, 2))
	fmt.Println(s[3](10, 2))
}
