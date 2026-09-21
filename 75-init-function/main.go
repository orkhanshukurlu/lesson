package main

import "fmt"

var m = map[string]string{}

func init() {
	fmt.Println("init function called")

	m["a"] = "A"
	m["b"] = "B"
}

func main() {
	fmt.Println("main function called")
	fmt.Println(m)
}
