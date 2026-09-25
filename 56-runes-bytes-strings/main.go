package main

import (
	"fmt"
)

func main() {
	a := "Jack"
	b := a[1]
	fmt.Println(b)        // 97
	fmt.Printf("%T\n", b) // uint8

	c := "café"
	fmt.Println(len(c)) // 5

	d := "é"
	fmt.Println(d[0]) // 195
	for _, r := range d {
		fmt.Println(r) // 233
	}
}
