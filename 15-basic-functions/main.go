package main

import "fmt"

func main() {
	hello("World") // Hello, World!
}

func hello(name string) {
	fmt.Println("Hello,", name)
}
