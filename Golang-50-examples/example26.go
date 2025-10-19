// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	for i, v := range arr {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

