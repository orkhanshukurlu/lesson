package main

import "fmt"

func main() {
	b := 0

	for {
		if b == 3 {
			break
		}

		fmt.Println(b) // 0 1 2
		b++
	}

	for c := range 3 {
		if c == 1 {
			continue
		}

		fmt.Println(c) // 0 2
	}

	for i := range 3 {
		for j := range 3 {
			if i == 1 {
				break
			}

			fmt.Println(i, j)
			/*
				0 0
				0 1
				0 2
				2 0
				2 1
				2 2
			*/
		}
	}
}
