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
	fmt.Println(s[1]) // 9

	a := []int{1, 2}
	b := a

	a[1] = 19
	fmt.Println(a[1]) // 19
	fmt.Println(b[1]) // 19

	c := []int{1, 2, 3}
	reset(c)
	fmt.Println(c[0]) // 1

	d := []int{1, 2}
	e := d
	d = append(d, 3)
	d[1] = 19

	fmt.Println(d) // [1 19 3]
	fmt.Println(e) // [1 2]

	m := make([]int, 2)
	m[0], m[1] = 1, 2

	m2 := addValue(m, 20)
	fmt.Println(m)  // [1 2]
	fmt.Println(m2) // [1 2 20]
}
