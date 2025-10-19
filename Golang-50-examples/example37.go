// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Defer
	defer fmt.Println("This will print last")
	fmt.Println("This will print first")
	fmt.Println("This will print second")
}
