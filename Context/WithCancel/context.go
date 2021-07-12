package main

import (
	"context"
	"fmt"
	"time"
)

func RunWithContext(ctx context.Context) {
	t := time.After(2 * time.Second)

	select {
	case <-t:
		fmt.Println("Doing thing I was supposed to do")
	case <-ctx.Done():
		fmt.Println("Context has finished!")
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		RunWithContext(ctx)
	}()

	time.Sleep(1 * time.Second)

	// Simulate external cancel event
	fmt.Println("Sending Cancel")
	cancel()

	fmt.Println("Context error: ", ctx.Err())
	fmt.Print("Context deadline: ")
	fmt.Println(ctx.Deadline())

	fmt.Println("Done")
}

func main() {
	Run()

	// Show cleanup
	time.Sleep(100 * time.Millisecond)
}
