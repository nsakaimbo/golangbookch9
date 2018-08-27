package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {

	go func() {
		var mu sync.Mutex

		mu.Lock()
		fmt.Println("first lock")

		// Uncomment to deadlock
		// mu.Lock()
		// fmt.Println("2nd lock")
		// mu.Unlock()

		mu.Unlock()
		fmt.Println("exiting goroutine")
	}()

	time.Sleep(5 * time.Second) // wait 5 seconds
}

// END OMIT
