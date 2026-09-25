package main

import "fmt"

func makeCounter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {
	count := 0

	increment := func() {
		fmt.Println(count)
		count++
	}

	increment() // 0
	increment() // 1
	increment() // 2

	counter := makeCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
}
