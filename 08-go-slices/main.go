package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)    // [1 2 3]
	fmt.Println(x[0]) // 1

	x = append(x, 4, 5)
	fmt.Println(x[4])   // 5
	fmt.Println(len(x)) // 5

	y := make([]string, 2)
	y[0] = "Hello"
	y[1] = "World"
	fmt.Println(y) // [Hello World]

	z := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := z[1:]
	t2 := z[:4]
	t3 := z[2:5]
	fmt.Println(t1) // [2 3 4 5 6 7]
	fmt.Println(t2) // [1 2 3 4]
	fmt.Println(t3) // [3 4 5]
}
