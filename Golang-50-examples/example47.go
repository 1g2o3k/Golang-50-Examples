// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"sync"
)

func main() {
	// Example: Mutex for safe concurrent access
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
