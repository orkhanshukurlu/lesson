package main

import "fmt"

func score() (int, int) {
	return 1, 2
}

func main() {
	a, b := 1, 2
	fmt.Println("a:", a, "b:", b)

	c, d := score()
	fmt.Println("c:", c, "d:", d)

	e, f := 1, 2
	e, f = f, e
	fmt.Println("e:", e, "f:", f)
}
