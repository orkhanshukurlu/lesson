package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "alice@example.com"
	fmt.Println(strings.Contains(email, "@"))
	fmt.Println(strings.HasPrefix(email, "alice"))
	fmt.Println(strings.HasSuffix(email, ".com"))

	str1 := "   Hello   "
	fmt.Println(strings.TrimSpace(str1))
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.ToLower(str1))

	str2 := "Hello John John John"
	fmt.Println(strings.Replace(str2, "John", "Alice", 2))
	fmt.Println(strings.Replace(str2, "John", "Alice", -1))
	fmt.Println(strings.ReplaceAll(str2, "John", "Alice"))
}
