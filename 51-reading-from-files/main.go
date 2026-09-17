package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err1 := os.Open("output.txt")
	if err1 != nil {
		fmt.Println("Error opening file:", err1)
		return
	}
	defer file.Close()

	content, err2 := io.ReadAll(file)
	if err2 != nil {
		fmt.Println("Error reading file:", err2)
		return
	}

	fmt.Println(string(content))
}
