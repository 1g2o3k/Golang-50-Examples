// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	for i := 0; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("is not divided into two")
		}
	}
}

