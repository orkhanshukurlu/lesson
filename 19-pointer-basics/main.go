package main

import "fmt"

func main() {
	a := 987
	b := &a

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("*b: ", *b)

	*b = 123
	fmt.Println("a: ", a)
}
