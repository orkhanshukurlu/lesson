package main

import "fmt"

type User struct {
	Name  string
	Score int
}

func addTen(x *int) {
	*x += 10
}

func incrementScore(u *User) {
	u.Score++
}

func main() {
	x := 5
	addTen(&x)
	fmt.Println(x) // 15

	u := User{Name: "John", Score: 5}
	incrementScore(&u)
	fmt.Println(u.Score) // 6
}
