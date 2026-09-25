package main

import "fmt"

func divide(a, b int) (int, int, bool) {
	if b == 0 {
		return 0, 0, false
	}

	return a / b, a % b, true
}

func main() {
	a, b, c := divide(10, 2)

	if !c {
		fmt.Println("Cannot divide by zero")
	} else {
		fmt.Println(a) // 5
		fmt.Println(b) // 0
	}
}
