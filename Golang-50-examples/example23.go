// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func monday() {
	fmt.Println("Monday")
}

func tuesday() {
	fmt.Println("Tuesday")
}

func wednesday() {
	fmt.Println("Wednesday")
}

func thursday() {
	fmt.Println("Thursday")
}

func friday() {
	fmt.Println("Friday")
}

func saturday() {
	fmt.Println("Saturday")
}

func sunday() {
	fmt.Println("Sunday")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter day number: ")
	dayInput, _ := reader.ReadString('\n')
	dayNumber, err := strconv.Atoi(strings.TrimSpace(dayInput))
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	switch dayNumber {
	case 1:
		monday()
	case 2:
		tuesday()
	case 3:
		wednesday()
	case 4:
		thursday()
	case 5:
		friday()
	case 6:
		saturday()
	case 7:
		sunday()
	default:
		fmt.Println("Invalid day number. Please enter 1-7.")
	}
}

