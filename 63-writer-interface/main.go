package main

import (
	"fmt"
	"io"
	"os"
)

func writeMessage(w io.Writer) {
	_, err := w.Write([]byte("Hello, World!\n"))

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	writeMessage(os.Stdout)

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writeMessage(file)
}
