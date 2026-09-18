package main

import "fmt"

type UserID int

func main() {
	var id UserID = 4
	fmt.Println(id)
	fmt.Printf("%T\n", id)

	num := 10
	n1 := UserID(num)
	n2 := int(id)
	fmt.Println(n1, n2)
}
