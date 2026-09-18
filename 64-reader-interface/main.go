package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readMessage(r io.Reader) {
	msg, err := io.ReadAll(r)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(msg))
}

func main() {
	reader := strings.NewReader("Hello, World!")
	readMessage(reader)

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	readMessage(file)
}
