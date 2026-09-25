package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	field string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for field %s", e.field)
}

func handleError(s string) error {
	if s == "" {
		return ValidationError{field: "name"}
	}
	return nil
}

func main() {
	s := errors.New("error")
	fmt.Println(s) // error

	err := handleError("")
	if err != nil {
		fmt.Println(err) // validation error for field name
	} else {
		fmt.Println("no error")
	}
}
