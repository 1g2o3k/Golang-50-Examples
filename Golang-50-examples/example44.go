// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Example: Regular expressions
	text := "Hello, my email is user@example.com and my phone is 123-456-7890"
	emailRegex := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	phoneRegex := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)

	emails := emailRegex.FindAllString(text, -1)
	phones := phoneRegex.FindAllString(text, -1)

	fmt.Println("Emails found:", emails)
	fmt.Println("Phones found:", phones)

	// Replace example
	replaced := emailRegex.ReplaceAllString(text, "[EMAIL]")
	fmt.Println("Text with emails replaced:", replaced)
}
