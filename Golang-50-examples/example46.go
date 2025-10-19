// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Example: Goroutines with sync.WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed")
}
