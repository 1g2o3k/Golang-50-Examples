// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	var number int
	fmt.Println("day number:")
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println("Error:", err)
	}
	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}
}

