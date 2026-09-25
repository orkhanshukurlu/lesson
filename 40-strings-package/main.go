package main

import (
	"fmt"
	"strings"
)

func main() {
	email := "alice@example.com"
	fmt.Println(strings.Contains(email, "@"))      // true
	fmt.Println(strings.HasPrefix(email, "alice")) // true
	fmt.Println(strings.HasSuffix(email, ".com"))  // true

	str1 := "   Hello   "
	fmt.Println(strings.TrimSpace(str1)) // Hello
	fmt.Println(strings.ToUpper(str1))   //    HELLO
	fmt.Println(strings.ToLower(str1))   //    hello

	str2 := "Hello John John John"
	fmt.Println(strings.Replace(str2, "John", "Alice", 2))  // Hello Alice Alice John
	fmt.Println(strings.Replace(str2, "John", "Alice", -1)) // Hello Alice Alice Alice
	fmt.Println(strings.ReplaceAll(str2, "John", "Alice"))  // Hello Alice Alice Alice
}
