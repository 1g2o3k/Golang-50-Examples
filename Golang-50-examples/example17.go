// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	size := 12
	fmt.Printf("=== Multiplication Table (%dx%d) ===\n", size, size)
	fmt.Printf("%4s", "×")
	for i := 1; i <= size; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()
	fmt.Printf("%4s", "---")
	for i := 1; i <= size; i++ {
		fmt.Printf("%4s", "----")
	}
	fmt.Println()
	for i := 1; i <= size; i++ {
		fmt.Printf("%2d |", i)
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

