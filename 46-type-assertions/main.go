package main

import "fmt"

func main() {
	var x any = "Hello"

	if s, ok := x.(int); !ok {
		fmt.Println("Not a string")
	} else {
		fmt.Println(s)
	}

	switch x.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown type")
	}

	t := x.(int)
	fmt.Println(t)
}
