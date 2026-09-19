package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Hello from main!")

	go func() {
		fmt.Println("Hello from anonymous goroutine!")
	}()

	for range 1000000 {
	}

	fmt.Println("Hello from main!")
}
