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

func main() {
	admin1 := admin{
		name:     "John",
		surname:  "Doe",
		password: "1234",
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name, surname and password (space separated):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	// Boşluklara göre ayır
	parts := strings.Split(input, " ")
	if len(parts) < 3 {
		fmt.Println("Error: Please enter name, surname and password")
		return
	}
	inputName := parts[0]
	inputSurname := parts[1]
	inputPassword := parts[2]
	if inputName == admin1.name && inputSurname == admin1.surname && inputPassword == admin1.password {
		fmt.Println("Access granted")
		fmt.Printf("Welcome %s %s!\n", admin1.name, admin1.surname)
	} else {
		fmt.Println("Access denied")
		fmt.Println("Invalid credentials")
	}
}

