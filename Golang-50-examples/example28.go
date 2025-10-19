// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Using slices
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println("Slice:", slice)
	slice = append(slice, 60)
	fmt.Println("After append:", slice)
}
