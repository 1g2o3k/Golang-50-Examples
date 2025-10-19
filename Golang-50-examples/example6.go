// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	var name string
	var surname string
	var year int
	fmt.Println("Your name, surname and year:")
	user, err := fmt.Scanf("%s %s %d", &name, &surname, &year)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Successfully read %d values\n", user)
	fmt.Println("Your name is", name, "your surname is", surname, "your year is", year)
}

