// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Println("Enter a number:")
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	numStr := strconv.Itoa(input)
	fmt.Println("You entered number:", numStr)
}

