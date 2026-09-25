package main

import "fmt"

type User struct {
	Name string
}

func main() {
	u1 := new(User)
	u1.Name = "John"
	fmt.Println(u1) // &{John}

	u2 := User{Name: "Jane"}
	fmt.Println(u2) // {Jane}
}
