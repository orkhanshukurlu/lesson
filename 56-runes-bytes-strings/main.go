package main

import (
	"fmt"
)

func main() {
	a := "Jack"
	b := a[1]
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	c := "café"
	fmt.Println("Length:", len(c))

	d := "é"
	fmt.Println(d[0])
	for _, r := range d {
		fmt.Println(r)
	}
}
