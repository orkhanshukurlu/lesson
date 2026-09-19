package main

import "fmt"

func main() {
	b := 0

	for {
		if b == 3 {
			break
		}

		fmt.Println(b)
		b++
	}

	for c := range 3 {
		if c == 1 {
			continue
		}

		fmt.Println(c)
	}

	for i := range 3 {
		for j := range 3 {
			if i == 1 {
				break
			}

			fmt.Println(i, j)
		}
	}
}
