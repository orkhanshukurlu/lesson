package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}

	s := []string{"A", "B", "C"}

	for key, value := range s {
		println(key, value)
	}
}
