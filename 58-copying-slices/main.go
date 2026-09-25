package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 5
	fmt.Println(a) // [5 2 3]
	fmt.Println(b) // [5 2 3]

	source := []int{1, 2, 3}
	target := make([]int, len(source))
	copy(target, source)
	target[0] = 5
	fmt.Println(source) // [1 2 3]
	fmt.Println(target) // [5 2 3]

	c := []int{1, 2, 3, 4}
	d := make([]int, 2)
	copy(d, c)
	fmt.Println(d) // [1 2]

	e := []int{1, 2, 3, 4}
	copy(e[1:], e)
	fmt.Println(e) // [1 1 2 3]
}
