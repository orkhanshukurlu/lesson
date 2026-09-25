package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num) // 123
	}

	n := 123
	s := strconv.Itoa(n)
	fmt.Println(s) // "123"

	parseInt, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(parseInt) // 123
	}
}
