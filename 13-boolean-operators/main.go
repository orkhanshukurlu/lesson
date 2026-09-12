package main

import (
	"fmt"
	"math/rand"
)

func main() {
	isMember := rand.Intn(2) == 1
	age := rand.Intn(50) + 1
	hasPermission := rand.Intn(2) == 1

	if isMember && (age > 18 || hasPermission) {
		fmt.Println("You can enter the club!")
	} else {
		fmt.Println("You cannot enter the club!")
	}
}
