// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"time"
)

func main() {
	// Example: Time and date handling
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())

	// Formatting time
	fmt.Println("Formatted time:", now.Format("2006-01-02 15:04:05"))

	// Parsing time
	layout := "2006-01-02"
	dateStr := "2023-12-25"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", parsedTime)
}
