package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	fmt.Println("First element:", a[0])
	a[0] = 4
	fmt.Println("First element:", a[0])
	fmt.Println("Array length:", len(a))

	b := [2]string{"Hello", "World"}
	fmt.Println(b)
}
