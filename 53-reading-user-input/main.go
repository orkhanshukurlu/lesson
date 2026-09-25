package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter two numbers: ") // 5 4

	count, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(a + b) // 9
	fmt.Println(count) // 2
}
