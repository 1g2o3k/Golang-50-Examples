// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Example: Reflection
	x := 42
	v := reflect.ValueOf(x)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Value:", v.Int())

	// Reflecting on a struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	v2 := reflect.ValueOf(p)
	t := v2.Type()
	for i := 0; i < v2.NumField(); i++ {
		field := v2.Field(i)
		fieldType := t.Field(i)
		fmt.Printf("Field %s (%s): %v\n", fieldType.Name, fieldType.Type, field.Interface())
	}
}
