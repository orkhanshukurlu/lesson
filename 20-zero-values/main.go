package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool
	var d float64
	var e complex64
	var f byte
	var g rune

	type S struct {
		Name string
	}

	var s S

	fmt.Println("int: ", a)
	fmt.Println("string: ", b)
	fmt.Println("bool: ", c)
	fmt.Println("float64: ", d)
	fmt.Println("complex64: ", e)
	fmt.Println("byte: ", f)
	fmt.Println("rune: ", g)
	fmt.Println("struct: ", s.Name)
}
