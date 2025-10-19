// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func fibonacciIterative(n int) {
	a, b := 0, 1
	fmt.Printf("Fibonacci (%d): ", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	fibonacciIterative(100)
}

