// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i string
	fmt.Println("Input:")
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		fmt.Println("Scan Error:", err)
		return
	}
	numint, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println("Conversion Error:", err)
		return
	}
	fmt.Println("Converted number:", numint)
}

