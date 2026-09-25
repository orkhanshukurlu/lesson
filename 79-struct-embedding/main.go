package main

import (
	"fmt"
)

type Address struct {
	City   string
	Street string
}

type User struct {
	Name string
	Address
}

func (a *Address) Full() string {
	return fmt.Sprintf("%s, %s", a.City, a.Street)
}

func main() {
	u := User{
		Name:   "John Doe",
		City:   "New York",
		Street: "123 Main St",
	}

	fmt.Println(u)        // {John Doe {New York 123 Main St}}
	fmt.Println(u.Full()) // New York, 123 Main St
}
