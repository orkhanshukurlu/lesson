package main

import (
	"fmt"
	"io"
	"os"
)

func writeMessage(w io.Writer) {
	_, err := w.Write([]byte("Hello, World!\n"))

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	writeMessage(os.Stdout)

	file, err := os.Create("output.txt")
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

	writeMessage(file) // Hello, World!
}
