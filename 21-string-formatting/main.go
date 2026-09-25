package main

import "fmt"

func main() {
	fmt.Printf("I'm %s, my age is %d\n", "John", 30) // I'm John, my age is 30
	s := fmt.Sprintf("I'm %s", "John")
	fmt.Println(s) // I'm John
}
