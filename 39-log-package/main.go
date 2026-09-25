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
		log.Fatal(err)
		/*
			2026/09/25 15:33:39 division by zero
			exit status 1
		*/
	}
	log.Println(r)
}
