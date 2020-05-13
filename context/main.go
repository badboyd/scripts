package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer func() {
		fmt.Println("Cancel")
		cancel()
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

	time.Sleep(2 * time.Second)
	fmt.Println("finish")

	fmt.Println(time.Duration(35) * time.Second)

	t, _ := time.ParseDuration("35")
	fmt.Println(t)
	time.Sleep(t)
}
