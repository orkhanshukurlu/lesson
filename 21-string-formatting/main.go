package main

import "fmt"

func main() {
	fmt.Printf("I'm %s, my age is %d\n", "John", 30)
	s := fmt.Sprintf("I'm %s", "John")
	fmt.Println(s)
}
