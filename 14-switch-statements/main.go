package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	d := rand.IntN(10) + 1

	switch d {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown")
	}
}
