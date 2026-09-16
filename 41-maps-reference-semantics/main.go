package main

import "fmt"

func incrementMapValue(m map[string]int, key string, value int) {
	m[key] += value
}

func reset(m map[string]int) {
	m = map[string]int{}
}

func main() {
	m := map[string]int{
		"age": 32,
	}

	incrementMapValue(m, "age", 4)
	fmt.Println("Age:", m["age"])

	a := map[string]int{"age": 42}
	b := a

	a["age"] = 19
	fmt.Println("Age of a:", a["age"])
	fmt.Println("Age of b:", b["age"])

	c := map[string]int{"age": 21}
	reset(c)
	fmt.Println("Age of c:", c["age"])
}
