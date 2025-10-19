// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Error handling
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	result, _ := divide(10, 2)
	fmt.Println("Result:", result)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
