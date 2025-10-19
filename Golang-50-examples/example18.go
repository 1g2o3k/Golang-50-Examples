// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nEnter the size of multiplication table (1-20) or 'quit' to exit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "quit" {
			fmt.Println("Goodbye!")
			break
		}
		size, err := strconv.Atoi(input)
		if err != nil || size < 1 || size > 20 {
			fmt.Println("Please enter a valid number between 1 and 20")
			continue
		}
		printMultiplicationTable(size)
	}
}

func printMultiplicationTable(size int) {
	fmt.Printf("\n=== %dx%d Multiplication Table ===\n", size, size)
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

