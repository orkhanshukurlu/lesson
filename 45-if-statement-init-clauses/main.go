package main

import (
	"errors"
	"fmt"
)

func loadConfig() error {
	return errors.New("config not found")
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	if err := loadConfig(); err != nil {
		fmt.Println(err) // config not found
	} else {
		fmt.Println("config loaded")
	}

	if r, err := div(3, 0); err != nil {
		fmt.Println(err) // division by zero
	} else {
		fmt.Println(r)
	}
}
