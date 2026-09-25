package main

import "fmt"

func main() {
	var s []int
	var m map[string]string
	var p *int

	fmt.Println(s == nil) // true
	fmt.Println(m == nil) // true
	fmt.Println(p == nil) // true
}
