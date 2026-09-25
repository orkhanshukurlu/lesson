package main

import "fmt"

func info() (string, int) {
	return "John", 43
}

func main() {
	name, _ := info()
	fmt.Println(name) // John

	_, age := info()
	fmt.Println(age) // 43

	s := []string{"A", "B", "C"}

	for _, val := range s {
		fmt.Println(val) // A B C
	}
}
