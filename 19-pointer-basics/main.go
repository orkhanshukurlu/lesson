package main

import "fmt"

func main() {
	a := 987
	b := &a

	fmt.Println(a)  // 987
	fmt.Println(b)  // 0x3a0e4f81c008
	fmt.Println(*b) // 987

	*b = 123
	fmt.Println(a) // 123
}
