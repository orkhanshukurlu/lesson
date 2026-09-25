package main

import "fmt"

func score() (int, int) {
	return 1, 2
}

func main() {
	a, b := 1, 2
	fmt.Println(a, b) // 1 2

	c, d := score()
	fmt.Println(c, d) // 1 2

	e, f := 1, 2
	e, f = f, e
	fmt.Println(e, f) // 2 1
}
