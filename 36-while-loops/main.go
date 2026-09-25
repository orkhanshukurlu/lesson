package main

import "fmt"

func main() {
	n := 0
	for {
		if n == 3 {
			break
		}

		fmt.Println(n) // 0 1 2
		n++
	}

	m := 16
	for m > 0 {
		fmt.Println(m) // 16 8 4 2 1
		m /= 2
	}

	a := 0
	for ; a < 3; a++ {
		fmt.Println(a) // 0 1 2
	}

	b := 0
	for ; ; b++ {
		if b >= 3 {
			break
		}
		fmt.Println(b) // 0 1 2
	}
}
