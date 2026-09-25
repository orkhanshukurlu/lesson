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
	fmt.Println(n) // 25

	r, ok := div(10, 2)
	fmt.Println(r)  // 5
	fmt.Println(ok) // true
}
