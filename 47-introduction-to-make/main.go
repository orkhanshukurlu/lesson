package main

import "fmt"

func main() {
	s1 := make([]int, 3, 5)
	s2 := make([]bool, 2)

	fmt.Println("s1:", s1)
	fmt.Println("s1 len:", len(s1))
	fmt.Println("s1 cap:", cap(s1))

	fmt.Println("s2:", s2)
	fmt.Println("s2 len:", len(s2))
	fmt.Println("s2 cap:", cap(s2))

	m := make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20
	fmt.Println("m:", m)

	s3 := make([]int, 2)
	var s4 []int
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	s5 := make([]int, 2, 4)
	s5[0] = 1
	s5[1] = 2
	s5 = append(s5, 3, 4, 5, 6)
	fmt.Println("s5:", s5)
	fmt.Println("s5 len:", len(s5))
	fmt.Println("s5 cap:", cap(s5))
}
