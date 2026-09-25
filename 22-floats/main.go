package main

import "fmt"

func main() {
	var f1 float32 = 1.234
	var f2 = 1.234

	fmt.Println(f1)        // 1.234
	fmt.Println(f2)        // 1.234
	fmt.Printf("%T\n", f2) // float64

	var f3 float32 = 0.123456789
	var f4 = 0.123456789

	fmt.Println(f3) // 0.12345679
	fmt.Println(f4) // 0.123456789
}
