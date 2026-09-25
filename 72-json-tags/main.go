package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	p := Person{Name: "John", Age: 30}
	data, _ := json.Marshal(p)
	fmt.Println(string(data)) // {"name":"John","age":30}

	p2 := Person{Name: "Jane"}
	j, _ := json.Marshal(p2)
	fmt.Println(string(j)) // {"name":"Jane"}
}
