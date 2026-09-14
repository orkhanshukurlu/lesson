package main

import "fmt"

func describe(a string, b any) {
	fmt.Printf("%s: %v (%T)\n", a, b, b)
}

func main() {
	describe("Name", "Alice")
	describe("Age", 30)
	describe("Height", 1.75)
}
