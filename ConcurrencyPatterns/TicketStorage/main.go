package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Ticket storage
type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string // for simplicity: imagine this to be infinite
}

func (ts *TicketStore) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1             // Draw a ticket
	ts.slots[t] = s                                     // Store your data
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) { // Increase done
		runtime.Gosched()
	}
}

func (ts *TicketStore) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1] // Read up to done
}

var wg = sync.WaitGroup{}

func producer(id string, to_produce int, st *TicketStore) {
	for i := 1; i <= to_produce; i++ {
		// Simulate production effort
		time.Sleep(100 * time.Millisecond)

		fmt.Printf("[%s]Produced -> %d\n", id, i)
		st.Put(fmt.Sprintf("[%s]Stored %d", id, i))
	}

	wg.Done()
}

func main() {
	wg.Add(2)
	fmt.Println("Start producer/consumer")

	var store_ticket uint64
	var store_done uint64
	var store_slots []string = make([]string, 20)

	var store TicketStore = TicketStore{
		ticket: &store_ticket,
		done:   &store_done,
		slots:  store_slots,
	}

	go producer("1", 10, &store)
	go producer("2", 10, &store)

	wg.Wait()

	fmt.Println("Tickets done:")
	for _, slot := range store.slots {
		fmt.Println(slot)
	}

	fmt.Println("End producer/consumer")
}
