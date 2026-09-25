package main

import "fmt"

func main() {
	name := "John"
	greeting := "Hello " + name + "!"
	fmt.Println(greeting) // Hello John!

	str := `Hello John!
	My name is Alice`
	fmt.Println(str)
	/*
		Hello John!
		        My name is Alice
	*/

	name = "Alice"
	ascii := name[0]
	fmt.Println(name)          // Alice
	fmt.Println(ascii)         // 65
	fmt.Println(string(ascii)) // A
}
