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
		return -1, true
	}
}

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, ch chan<- int) {
	for i := 1; i <= to_produce; i++ {
		fmt.Printf("[%s]Produced -> %d\n", id, i)
		ch <- i
		// Simulate production effort
		time.Sleep(300 * time.Millisecond)
	}

	// Signal end of production
	close(ch)
	wg.Done()
}

func consumer(id string, ch <-chan int) {
	for {

		i, more := TryReceive(ch)

		fmt.Printf("[%s]More %t\n", id, more)

		if !more {
			break
		}

		fmt.Printf("[%s]Consumed <- %d\n", id, i)
		// Simulate consumption effort
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	fmt.Println("Start producer/consumer")

	chan_data := make(chan int, 10)

	go producer("1", 10, chan_data)
	go consumer("1", chan_data)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
