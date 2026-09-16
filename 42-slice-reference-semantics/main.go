package main

import "fmt"

func incrementSliceValue(s []int, value int) {
	s[1] += value
}

func reset(s []int) {
	s = []int{}
}

func addValue(s []int, value int) []int {
	return append(s, value)
}

func main() {
	s := []int{3, 5}

	incrementSliceValue(s, 4)
	fmt.Println("Second item:", s[1])

	a := []int{1, 2}
	b := a

	a[1] = 19
	fmt.Println("Second item of a:", a[1])
	fmt.Println("Second item of b:", b[1])

	c := []int{1, 2, 3}
	reset(c)
	fmt.Println("First item:", c[0])

	d := []int{1, 2}
	e := d
	d = append(d, 3)
	d[1] = 19

	fmt.Println("d:", d)
	fmt.Println("e:", e)

	m := make([]int, 2)
	m[0], m[1] = 1, 2

	m2 := addValue(m, 20)
	fmt.Println("Original slice after append:", m)
	fmt.Println("New slice after append:", m2)
}
