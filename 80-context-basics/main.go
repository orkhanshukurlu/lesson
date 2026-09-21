package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	<-ctx.Done()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel2()

	<-ctx2.Done()
	fmt.Println("Context timed out")

	ctx3, cancel3 := context.WithDeadline(
		context.Background(),
		time.Now().Add(2*time.Second),
	)

	defer cancel3()

	<-ctx3.Done()
	fmt.Println("Context deadline reached")
}
