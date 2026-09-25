package main

import "fmt"

func main() {
	x := 10
	y := float64(x)

	fmt.Println(y) // 10

	r := []rune{'a', 'b', 'c'}
	s := string(r)
	fmt.Println(s) // abc
}
