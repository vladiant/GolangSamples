package main

import (
	"fmt"
	"sync"
	"time"
)

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In { // Receive until closed
		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}

	// Signal end of dispatching
	close(OutA)
	close(OutB)
	wg.Done()
}

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, ch chan<- int) {
	for i := 1; i <= to_produce; i++ {
		fmt.Printf("[%s]Produced -> %d\n", id, i)
		ch <- i
		// Simulate production effort
		time.Sleep(100 * time.Millisecond)
	}

	// Signal end of production
	close(ch)
	wg.Done()
}

func consumer(id string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("[%s]Consumed <- %d\n", id, i)
		// Simulate consumption effort
		time.Sleep(300 * time.Millisecond)

	}
	wg.Done()
}

func main() {
	wg.Add(4)
	fmt.Println("Start producer/consumer")

	chan_in := make(chan int, 10)
	chan_out_a := make(chan int, 10)
	chan_out_b := make(chan int, 10)

	go producer("1", 10, chan_in)
	go Fanout(chan_in, chan_out_a, chan_out_b)
	go consumer("1", chan_out_a)
	go consumer("2", chan_out_b)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
