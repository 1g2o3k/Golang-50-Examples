// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// Example: JSON marshaling and unmarshaling
	p := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Marshal to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	// Unmarshal from JSON
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}
	fmt.Println("Unmarshaled:", p2)
}
