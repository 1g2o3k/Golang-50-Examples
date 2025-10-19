// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string
	var surname string
	var yearStr string
	fmt.Println("Your name, surname and year:")
	_, err := fmt.Scanf("%s %s %s", &name, &surname, &yearStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		fmt.Println("Conversion Error:", err)
		return
	}
	if year < 18 {
		fmt.Println("You must be at least 18 years old.")
		return
	} else if year > 18 {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("ERROR")
	}
}

