// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"os"
)

func main() {
	// Example: File I/O - Reading from a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", count, data[:count])
}
