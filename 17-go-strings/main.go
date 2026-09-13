package main

import "fmt"

func main() {
	name := "John"
	greeting := "Hello " + name + "!"
	fmt.Println(greeting)

	str := `Hello John!
	My name is Alice`
	fmt.Println(str)

	name = "Alice"
	ascii := name[0]
	fmt.Println("String: ", name)
	fmt.Println("First ASCII: ", ascii)
	fmt.Println("First Character", string(ascii))
}
