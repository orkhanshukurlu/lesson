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

	fmt.Println(a)      // 0
	fmt.Println(b)      // ""
	fmt.Println(c)      // false
	fmt.Println(d)      // 0
	fmt.Println(e)      // (0+0i)
	fmt.Println(f)      // 0
	fmt.Println(g)      // 0
	fmt.Println(s.Name) // ""
}
