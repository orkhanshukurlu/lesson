package main

import (
	"fmt"
	"os"
)

func main() {
	file, err1 := os.Create("output.txt")
	if err1 != nil {
		fmt.Println("Error creating file:", err1)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully")
	fmt.Println("Filename:", file.Name())

	_, err2 := file.WriteString("1. Hello, world!\n")
	if err2 != nil {
		fmt.Println("Error writing to file:", err2)
	}
	_, err3 := file.WriteString("2. Hello, world!")
	if err3 != nil {
		fmt.Println("Error writing to file:", err3)
	}
}
