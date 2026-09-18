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
}
