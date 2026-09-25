package main

import (
	"fmt"
	"os"
)

func main() {
	file, err1 := os.Create("output.txt")
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

	fmt.Println("File created successfully")
	fmt.Println(file.Name()) // output.txt

	_, err2 := file.WriteString("1. Hello, world!\n")
	if err2 != nil {
		fmt.Println(err2)
	}
	_, err3 := file.WriteString("2. Hello, world!")
	if err3 != nil {
		fmt.Println(err3)
	}
}
