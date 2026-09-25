package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err1 := os.Open("output.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	content, err2 := io.ReadAll(file)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println(string(content))
	/*
		1. Hello, world!
		2. Hello, world!
	*/
}
