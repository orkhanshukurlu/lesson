package main

import "fmt"

func main() {
	x := 10
	y := float64(x)

	fmt.Println(y)

	r := []rune{'a', 'b', 'c'}
	s := string(r)
	fmt.Println(s)
}
