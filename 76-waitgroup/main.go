package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("Hello, World!")
	})

	wg.Wait()

	var wg2 sync.WaitGroup

	values := []string{"a", "b", "c"}

	for _, v := range values {
		wg2.Add(1)

		go func(t string) {
			defer wg2.Done()
			fmt.Println(t) // c b a
		}(v)
	}

	wg2.Wait()
}
