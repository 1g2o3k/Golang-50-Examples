// Donate: https://buymeacoffee.com/gokalp
package main

import "fmt"

// Example: Methods
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area of rectangle:", rect.Area())
}
