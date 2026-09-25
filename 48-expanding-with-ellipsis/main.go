package main

import "fmt"

func sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}
	return
}

func main() {
	values := []int{1, 2, 3}
	summed := sum(values...)
	fmt.Println(summed) // 6
}
