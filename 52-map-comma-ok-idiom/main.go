package main

import "fmt"

func main() {
	m := map[string]int{
		"sarah": 4,
		"bob":   5,
	}

	if val, ok := m["john"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("john not found")
	}

	fmt.Println(m["terry"])
}
