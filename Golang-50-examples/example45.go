// Donate: https://buymeacoffee.com/gokalp
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Example: Sorting slices
	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Println("Original slice:", numbers)

	sort.Ints(numbers)
	fmt.Println("Sorted slice:", numbers)

	// Sorting strings
	strings := []string{"banana", "apple", "cherry", "date"}
	fmt.Println("Original strings:", strings)

	sort.Strings(strings)
	fmt.Println("Sorted strings:", strings)

	// Custom sorting with sort.Slice
	people := []struct {
		name string
		age  int
	}{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	fmt.Println("Original people:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println("Sorted by age:", people)
}
