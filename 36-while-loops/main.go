package main

import "fmt"

func main() {
	n := 0
	for {
		if n == 3 {
			break
		}

		fmt.Println("n:", n)
		n++
	}

	m := 16
	for m > 0 {
		fmt.Println("m:", m)
		m /= 2
	}

	a := 0
	for ; a < 3; a++ {
		fmt.Println("a:", a)
	}

	b := 0
	for ; ; b++ {
		if b >= 3 {
			break
		}
		fmt.Println("b:", b)
	}
}
