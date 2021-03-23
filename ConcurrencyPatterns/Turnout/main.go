package main

import (
	"fmt"
	"sync"
	"time"
)

func Turnout(InA, InB <-chan int, OutA, OutB chan int) {
	defer wg.Done()

	for {
		var data int
		var more bool

		select { // Receive from first non-blocking channel
		case data, more = <-InA:
		case data, more = <-InB:
		}

		if !more {
			// ... ?
			fmt.Println("[Turnout] No more data")
			// Actually this is an antipattern ... (the receiver should never close)
			close(OutA)
			close(OutB)
			return
		}

		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
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
	fmt.Printf("[%s]Close channel\n", id)
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
	wg.Add(5)
	fmt.Println("Start producer/consumer")

	// Should Should not be buffered
	chan_in_a := make(chan int, 2)
	chan_in_b := make(chan int, 3)

	// Should not be buffered
	chan_out_a := make(chan int)
	chan_out_b := make(chan int)

	go producer("1", 10, chan_in_a)
	go producer("2", 10, chan_in_b)
	go Turnout(chan_in_a, chan_in_b, chan_out_a, chan_out_b)
	go consumer("1", chan_out_a)
	go consumer("2", chan_out_b)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
