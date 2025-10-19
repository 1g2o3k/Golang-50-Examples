// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type admin struct {
	name     string
	surname  string
	password string
}

var admins = []admin{
	{"John", "Doe", "1234"},
	{"Jane", "Smith", "admin"},
	{"Ahmet", "Yılmaz", "password"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("=== ADMIN LOGIN SYSTEM ===")
	fmt.Println("Enter 'quit' to exit")
	for {
		fmt.Println("\nPlease enter credentials:")
		fmt.Print("Name: ")
		inputName, _ := reader.ReadString('\n')
		inputName = strings.TrimSpace(inputName)
		if inputName == "quit" {
			fmt.Println("Goodbye!")
			break
		}
		fmt.Print("Surname: ")
		inputSurname, _ := reader.ReadString('\n')
		inputSurname = strings.TrimSpace(inputSurname)
		fmt.Print("Password: ")
		inputPassword, _ := reader.ReadString('\n')
		inputPassword = strings.TrimSpace(inputPassword)
		// Admin kontrolü
		authenticated := false
		var currentAdmin admin
		for _, admin := range admins {
			if inputName == admin.name && inputSurname == admin.surname && inputPassword == admin.password {
				authenticated = true
				currentAdmin = admin
				break
			}
		}
		if authenticated {
			fmt.Printf("\n Access granted! Welcome %s %s\n", currentAdmin.name, currentAdmin.surname)
			break
		} else {
			fmt.Println("\n Access denied! Invalid credentials")
			fmt.Println("Please try again or enter 'quit' to exit")
		}
	}
}

