// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Pointers
	x := 10
	ptr := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Value at pointer:", *ptr)
	*ptr = 20
	fmt.Println("New value of x:", x)
}
