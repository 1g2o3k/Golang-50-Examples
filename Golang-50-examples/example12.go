// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

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
	var inputName, inputSurname, inputPassword string
	fmt.Println("Enter your name, surname and password:")
	_, err := fmt.Scanf("%s %s %s", &inputName, &inputSurname, &inputPassword)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	if inputName == admin1.name && inputSurname == admin1.surname && inputPassword == admin1.password {
		fmt.Println("Access granted")
		fmt.Printf("Welcome %s %s!\n", admin1.name, admin1.surname)
	} else {
		fmt.Println("Access denied")
		fmt.Println("Invalid credentials")
	}
}

