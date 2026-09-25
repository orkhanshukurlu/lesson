package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0

	for range 5 {
		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			count++
		})
	}

	wg.Wait()
	fmt.Println(count) // 5
}
