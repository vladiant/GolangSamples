package main

import (
	"fmt"
	"sync"
	"time"
)

func TryReceive(c <-chan int) (data int, more bool) {
	select {
	case data, more = <-c:
		return data, more
	default: // processed when c is blocking
		return 0, true
	}
}

func Fanout(In <-chan int, OutA, OutB chan int) {
	for data := range In { // Receive until closed
		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}

	}
}

func TurnoutQuitChannel(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	defer wg.Done()

	for {
		var data int
		var more bool

		select { // Receive from first non-blocking channel
		case data, more = <-InA:
		case data, more = <-InB:

		case <-Quit: // Remember: Close generates a message
			fmt.Println("[TurnoutQuitChannel] Quit")
			close(InA) // Actually this is an antipattern ... (the receiver should never close)
			close(InB) // ... but you can argue this acts as a delegate

			Fanout(InA, OutA, OutB) // Flush the remaining data
			Fanout(InB, OutA, OutB)

			close(OutA)
			close(OutB)
			return
		}

		if !more {
			// ... ?
			fmt.Println("[TurnoutQuitChannel] No more data")
			return
		}

		select { // Send to first non-blocking channel
		case OutA <- data:
		case OutB <- data:
		}
	}
}

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, ch chan<- int, Quit <-chan int) {
	defer wg.Done()

	for i := 1; i <= to_produce; i++ {
		_, more := TryReceive(Quit)
		if !more {
			fmt.Printf("[%s]Stop signal received\n", id)
			return
		}

		fmt.Printf("[%s]Produced -> %d\n", id, i)
		ch <- i
		// Simulate production effort
		time.Sleep(100 * time.Millisecond)
	}
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
	wg.Add(6)
	fmt.Println("Start producer/consumer")

	chan_quit := make(chan int, 2)

	chan_in_a := make(chan int, 10)
	chan_in_b := make(chan int, 5)

	// At least one should not be buffered
	chan_out_a := make(chan int, 2)
	chan_out_b := make(chan int, 3)

	go producer("1", 10, chan_in_a, chan_quit)
	go producer("2", 10, chan_in_b, chan_quit)
	go TurnoutQuitChannel(chan_quit, chan_in_a, chan_in_b, chan_out_a, chan_out_b)

	go func(Quit chan<- int) {
		time.Sleep(400 * time.Millisecond)
		close(Quit)
		wg.Done()
	}(chan_quit)

	go consumer("1", chan_out_a)
	go consumer("2", chan_out_b)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
