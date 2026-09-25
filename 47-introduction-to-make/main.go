package main

import "fmt"

func main() {
	s1 := make([]int, 3, 5)
	s2 := make([]bool, 2)

	fmt.Println(s1)      // [0 0 0]
	fmt.Println(len(s1)) // 3
	fmt.Println(cap(s1)) // 5

	fmt.Println(s2)      // [false false]
	fmt.Println(len(s2)) // 2
	fmt.Println(cap(s2)) // 2

	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	fmt.Println(m) // map[key1:10 key2:20]

	s3 := make([]int, 2)
	var s4 []int
	fmt.Println(s3) // [0 0]
	fmt.Println(s4) // []

	s5 := make([]int, 2, 4)
	s5[0] = 1
	s5[1] = 2
	s5 = append(s5, 3, 4, 5, 6)
	fmt.Println(s5)      // [1 2 3 4 5 6]
	fmt.Println(len(s5)) // 6
	fmt.Println(cap(s5)) // 8
}
