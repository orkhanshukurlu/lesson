package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	age := rand.IntN(50) + 1

	if age >= 18 {
		fmt.Println("You are an adult")
	}
}
