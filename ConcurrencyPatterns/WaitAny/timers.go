package main

import (
	"fmt"
	"time"
)

func main() {
	// Creates a channel that will have something
	// added to it after the given duration
	t1 := time.After(1 * time.Second)
	t2 := time.After(2 * time.Second)

	select {
	case <-t1:
		fmt.Println("Timer 1 fired!")
	case <-t2:
		fmt.Println("Timer 2 fired!")
	}

	fmt.Println("Done")
}
