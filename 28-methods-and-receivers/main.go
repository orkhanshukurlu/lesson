package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func main() {
	u := User{FirstName: "John", LastName: "Doe"}
	fmt.Println(u.FullName()) // John Doe
}
