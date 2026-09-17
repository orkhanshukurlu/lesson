package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Weekday:", now.Weekday())
	fmt.Println("Location:", now.Location())

	formatted := now.Format("2006-01-02 15:04")
	fmt.Println("Formatted time:", formatted)
}
