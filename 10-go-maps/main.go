package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(m) // map[a:1 b:2 c:3]

	m["a"] = 5
	m["d"] = 4
	fmt.Println(m["a"], m["d"]) // 5 4

	val, exits := m["x"]
	fmt.Println(val, exits) // 0 false

	delete(m, "b")
	fmt.Println(m) // map[a:5 c:3 d:4]

	n := make(map[string]int)
	n["x"] = 10
	fmt.Println(n) // map[x:10]
}
