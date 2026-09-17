package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Enter two numbers: ")

	count, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	fmt.Println("The sum of the two numbers is:", a+b)
	fmt.Println("The number of arguments read:", count)
}
