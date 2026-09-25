package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // something bad happened
		}
	}()

	panic("something bad happened")
}
