package main

// https://gist.github.com/montanaflynn/d9b358249939c5c541ec
// https://gobyexample.com/select

import (
	"fmt"
	"time"
)

func waitForChannelsToClose(chans ...<-chan time.Time) {
	t := time.Now()
	for _, v := range chans {
		<-v
		fmt.Printf("%v for chan to close\n", time.Since(t))
	}
	fmt.Printf("%v for channels to close\n", time.Since(t))
}

func main() {
	// Creates a channel that will have something
	// added to it after the given duration
	t1 := time.After(1 * time.Second)
	t2 := time.After(2 * time.Second)

	waitForChannelsToClose(t1, t2)
	fmt.Println("Done")
}
