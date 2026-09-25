package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func greet(prefix string, names ...string) {
	for _, name := range names {
		fmt.Println(prefix, name)
	}
}

func main() {
	a := sum(1, 2, 3)
	b := sum(4, 5, 6)
	c := sum()

	fmt.Println(a, b, c) // 6 15 0

	greet("Hello", "Alice", "Bob", "Charlie")
	/*
		Hello Alice
		Hello Bob
		Hello Charlie
	*/
}
