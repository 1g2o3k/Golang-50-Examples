// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	// Example: Structs
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Person:", p)
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
