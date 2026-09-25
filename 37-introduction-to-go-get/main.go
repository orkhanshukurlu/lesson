package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println(id) // 531806df-d88b-41b5-a9de-10e2c6d67160
}
