package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	fmt.Println("First element:", x[0])

	x = append(x, 4, 5)
	fmt.Println("Last element:", x[4])
	fmt.Println("Slice length:", len(x))

	y := make([]string, 2)
	y[0] = "Hello"
	y[1] = "World"
	fmt.Println(y)

	z := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := z[1:]
	t2 := z[:4]
	t3 := z[2:5]
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)
}
