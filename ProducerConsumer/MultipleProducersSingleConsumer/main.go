// https://medium.com/@matryer/golang-advent-calendar-day-two-starting-and-stopping-things-with-a-signal-channel-f5048161018

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, ch chan<- int, ch_done chan<- struct{}) {
	for i := 1; i <= to_produce; i++ {
		fmt.Printf("[%s]Produced -> %d\n", id, i)
		ch <- i
		// Simulate production effort
		time.Sleep(20 * time.Millisecond)
	}

	// Signal end of production
	ch_done <- struct{}{}

	wg.Done()
}

// Listens to close signals from registered number of producers
// After all registered producers sent close signal
// the communication channel is closed
func listener(count int, ch_done <-chan struct{}, ch chan<- int) {
	for received := 0; received < count; received++ {
		// <-ch_done is sufficient
		_, ok := <-ch_done
		fmt.Println("Producer closed")
		if !ok {
			break
		}
	}

	fmt.Println("Channel closed")
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
		time.Sleep(100 * time.Millisecond)

	}
	wg.Done()
}

func main() {
	wg.Add(4)
	fmt.Println("Start producer/consumer")

	chan_data := make(chan int, 10)
	chan_done := make(chan struct{})

	go listener(2, chan_done, chan_data)
	go producer("1", 10, chan_data, chan_done)
	go producer("2", 12, chan_data, chan_done)
	go consumer("1", chan_data)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
