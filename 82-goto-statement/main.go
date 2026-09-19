package main

import "fmt"

func main() {
	fmt.Println("Start of the program")
	goto MyLabel

MyLabel:
	fmt.Println("Stepped inside the label")
}
