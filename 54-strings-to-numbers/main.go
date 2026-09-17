package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		fmt.Println("Converted integer:", num)
	}

	n := 123
	s := strconv.Itoa(n)
	fmt.Println("Converted string:", s)

	parseInt, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		fmt.Println("Converted integer:", parseInt)
	}
}
