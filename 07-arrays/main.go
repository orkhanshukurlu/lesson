package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)    // [1 2 3]
	fmt.Println(a[0]) // 1
	a[0] = 4
	fmt.Println(a[0])   // 4
	fmt.Println(len(a)) // 3

	b := [2]string{"Hello", "World"}
	fmt.Println(b) // [Hello World]
}
