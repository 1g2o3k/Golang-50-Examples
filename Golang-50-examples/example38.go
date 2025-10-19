// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Panic and recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went wrong!")
}
