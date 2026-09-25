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
		fmt.Println(err)
	}

	fmt.Println(string(msg))
}

func main() {
	reader := strings.NewReader("Hello, World!")
	readMessage(reader) // Hello, World!

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	readMessage(file) // Hello, World!
}
