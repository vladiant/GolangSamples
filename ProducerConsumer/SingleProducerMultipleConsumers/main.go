package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, ch chan<- int) {
	for i := 1; i <= to_produce; i++ {
		fmt.Printf("[%s]Produced -> %d\n", id, i)
		ch <- i
		// Simulate production effort
		time.Sleep(200 * time.Millisecond)
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
		time.Sleep(100 * time.Millisecond)

	}
	wg.Done()
}

func main() {
	wg.Add(3)
	fmt.Println("Start producer/consumer")

	chan_data := make(chan int)

	go producer("1", 10, chan_data)
	go consumer("1", chan_data)
	go consumer("2", chan_data)

	wg.Wait()
	fmt.Println("End producer/consumer")
}
