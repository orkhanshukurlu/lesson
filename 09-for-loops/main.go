package main

import "fmt"

func main() {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i) // 0 2 4 6 8
	}

	s := []string{"A", "B", "C"}

	for key, value := range s {
		fmt.Println(key, value) // 0 A 1 B 2 C
	}
}
