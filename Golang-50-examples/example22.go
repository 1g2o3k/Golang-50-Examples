// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getScore(reader *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strconv.ParseFloat(strings.TrimSpace(input), 64)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scores := make([]float64, 3)
	var err error
	for i := 0; i < 3; i++ {
		scores[i], err = getScore(reader, fmt.Sprintf("Score %d: ", i+1))
		if err != nil {
			fmt.Println("Invalid input. Please enter numeric values.")
			return
		}
	}
	average := (scores[0] + scores[1] + scores[2]) / 3
	if average >= 50 {
		fmt.Printf("Passed - Average: %.2f\n", average)
	} else {
		fmt.Printf("Failed - Average: %.2f\n", average)
	}
}

