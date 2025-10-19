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
	fmt.Println("score 1:")
	not1, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("score 2:")
	not2, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println("score 3:")
	not3, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	//data type conversion
	not1f, err := strconv.ParseFloat(strings.TrimSpace(not1), 64)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	not2f, err := strconv.ParseFloat(strings.TrimSpace(not2), 64)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	not3f, err := strconv.ParseFloat(strings.TrimSpace(not3), 64)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	average := (not1f + not2f + not3f) / 3
	if average > 50 {
		fmt.Printf("Passed with average: %.2f\n", average)
	} else {
		fmt.Printf("Failed with average: %.2f\n", average)
	}
}

