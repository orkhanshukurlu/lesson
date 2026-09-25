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
	/*
		panic: interface conversion: interface {} is string, not int

		goroutine 1 [running]:
		main.main()
		        /home/orkhan/Projects/lesson/46-type-assertions/main.go:23 +0x9f
		exit status 2
	*/
}
