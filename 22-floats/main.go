package main

import "fmt"

func main() {
	var f1 float32 = 1.234
	var f2 = 1.234

	fmt.Println("f1: ", f1)
	fmt.Println("f2: ", f2)
	fmt.Printf("default type: %T\n", f2)

	var f3 float32 = 0.123456789
	var f4 = 0.123456789

	fmt.Println("f3: ", f3)
	fmt.Println("f4: ", f4)
}
