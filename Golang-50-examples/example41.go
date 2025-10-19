// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"os"
)

func main() {
	// Example: File I/O - Writing to a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!\nThis is a test file.")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to file")
}
