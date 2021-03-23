package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Spinning CAS (Compare And Swap)
type Spinlock struct {
	state *int32
}

const free = int32(0)

func (l *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(l.state, free, 42) { // 42 or any other value but 0
		runtime.Gosched() // Poke the scheduler to prevent lockup
	}
}

func (l *Spinlock) Unlock() {
	atomic.StoreInt32(l.state, free) // Once atomic, always atomic
}

var wg = sync.WaitGroup{}

func worker(id string, iterations int, l *Spinlock) {
	for i := 1; i <= iterations; i++ {
		l.Lock()
		fmt.Printf("[%s]Iteration %d\n", id, i)
		// Simulate work effort
		time.Sleep(100 * time.Millisecond)
		l.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	fmt.Println("Start work")

	var lock_state int32
	lock := Spinlock{&lock_state}

	go worker("1", 10, &lock)
	go worker("2", 10, &lock)

	wg.Wait()
	fmt.Println("End work")
}
