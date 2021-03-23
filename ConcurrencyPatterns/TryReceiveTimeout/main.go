package main

import (
	"fmt"
	"sync"
	"time"
)

func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration): // time.After() returns a channel
		return -1, true, false
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

		i, more, ok := TryReceiveWithTimeout(ch, 50*time.Millisecond)

		fmt.Printf("[%s]More %t  OK: %t\n", id, more, ok)

		if !ok {
			fmt.Printf("[%s]Timeout!\n", id)
			continue
		}

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
