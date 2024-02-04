package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const gr = 100

	// declaring a WaitGroup to synchronize the goroutines with the main function.
	var wg sync.WaitGroup

	// adding 200 goroutines to the WaitGroup
	wg.Add(gr * 2)

	// Handling data race using mutex
	// 1.
	var m sync.Mutex
	// declaring a shared value
	var n int = 0

	// starting 200 goroutines
	for i := 0; i < gr; i++ {

		// each goroutine is an anonymous function
		go func() {

			// introducing some artificial time
			time.Sleep(time.Second / 10)
			m.Lock()
			// increment the shared value
			n++
			m.Unlock()
			// notifying the WaitGroup that the goroutine is done
			wg.Done()
		}()

		// goroutine that decrements the shared value
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			// m.Unlock()
			wg.Done()
		}()

	}
	// waiting for the goroutines to terminate.
	wg.Wait()

	//  printing the final value of n
	fmt.Println("The final value if n: ", n) // it will have another value for each program execution -> DATA RACE
}
