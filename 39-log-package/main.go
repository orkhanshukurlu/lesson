package main

import (
	"errors"
	"log"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	r, err := div(10, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Println("Result:", r)
}
