package main

import "fmt"

func describe(a string, b any) {
	fmt.Printf("%s: %v (%T)\n", a, b, b)
}

func main() {
	describe("Name", "Alice") // Name: Alice (string)
	describe("Age", 30)       // Age: 30 (int)
	describe("Height", 1.75)  // Height: 1.75 (float64)
}
