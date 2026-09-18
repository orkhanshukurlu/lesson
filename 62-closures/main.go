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

	increment()
	increment()
	increment()

	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
