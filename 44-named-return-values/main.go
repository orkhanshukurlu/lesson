package main

import "fmt"

func square(n int) (result int) {
	result = n * n
	return
}

func div(a, b int) (d int, ok bool) {
	if b != 0 {
		d = a / b
		ok = true
	} else {
		d = 0
		ok = false
	}
	return
}

func main() {
	n := square(5)
	fmt.Println("square of 5:", n)

	r, ok := div(10, 2)
	fmt.Println("r:", r)
	fmt.Println("ok:", ok)
}
