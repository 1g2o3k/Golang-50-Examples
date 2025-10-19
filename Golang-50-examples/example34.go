// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Goroutines
	go func() {
		fmt.Println("Hello from goroutine!")
	}()
	fmt.Println("Hello from main!")
	// Note: In a real program, you'd use sync.WaitGroup or similar to wait for goroutines
}
