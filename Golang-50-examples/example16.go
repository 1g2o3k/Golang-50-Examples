// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	fmt.Println("=== Multiplication Table (10x10) ===")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

