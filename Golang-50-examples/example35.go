// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Channels
	ch := make(chan string)
	go func() {
		ch <- "Hello from channel!"
	}()
	msg := <-ch
	fmt.Println(msg)
}
