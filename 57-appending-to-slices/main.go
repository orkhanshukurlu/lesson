package main

import "fmt"

func main() {
	s := make([]int, 0, 2)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(len(s)) // 2
	fmt.Println(cap(s)) // 2
	s = append(s, 3)
	fmt.Println(len(s)) // 3
	fmt.Println(cap(s)) // 4
}
