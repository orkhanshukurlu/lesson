package main

import "fmt"

func main() {
	msg := func(name string) {
		fmt.Println("Hello:", name) // Hello: John
	}

	msg("John")

	func(name string) {
		fmt.Println("Hello:", name) // Hello: Alice
	}("Alice")

	var m func(string)
	m = func(name string) {
		fmt.Println("Hello:", name) // Hello: Bob
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

	fmt.Println(s[0](10, 2)) // 12
	fmt.Println(s[1](10, 2)) // 8
	fmt.Println(s[2](10, 2)) // 20
	fmt.Println(s[3](10, 2)) // 5
}
