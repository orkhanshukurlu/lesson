package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)            // 2026-09-25 16:17:15.746150002 +0400 +04 m=+0.000012750
	fmt.Println(now.Year())     // 2026
	fmt.Println(now.Month())    // September
	fmt.Println(now.Day())      // 25
	fmt.Println(now.Hour())     // 16
	fmt.Println(now.Minute())   // 17
	fmt.Println(now.Second())   // 15
	fmt.Println(now.Weekday())  // Friday
	fmt.Println(now.Location()) // Local

	formatted := now.Format("2006-01-02 15:04")
	fmt.Println(formatted) // 2026-09-25 16:17
}
