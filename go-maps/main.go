package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(m)

	m["a"] = 5
	m["d"] = 4
	fmt.Println(m["a"], m["d"])

	val, exits := m["x"]
	fmt.Println(val, exits)

	delete(m, "b")
	fmt.Println(m)

	n := make(map[string]int)
	n["x"] = 10
	fmt.Println(n)
}
