package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[0] = 5
	fmt.Println(a)
	fmt.Println(b)

	source := []int{1, 2, 3}
	target := make([]int, len(source))
	copy(target, source)
	target[0] = 5
	fmt.Println(source)
	fmt.Println(target)

	c := []int{1, 2, 3, 4}
	d := make([]int, 2)
	copy(d, c)
	fmt.Println(d)

	e := []int{1, 2, 3, 4}
	copy(e[1:], e)
	fmt.Println(e)
}
