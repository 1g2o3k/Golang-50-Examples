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
	var total float64
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i <= 3; i++ {
		fmt.Printf("Score %d: ", i)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		score, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		total += score
	}
	average := total / 3
	if average > 50 {
		fmt.Printf("Passed with average: %.2f\n", average)
	} else {
		fmt.Printf("Failed with average: %.2f\n", average)
	}
}

