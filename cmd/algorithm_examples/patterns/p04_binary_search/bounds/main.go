package main

import "fmt"

// Запуск: go run ./bounds
func main() {
	fmt.Println("=== 04. Binary Search / Bounds (границы) ===")
	fmt.Println()

	sorted := []int{1, 3, 3, 5, 7, 7, 7, 9}
	fmt.Printf("массив: %v\n", sorted)
	fmt.Printf("LowerBound(_, 3)              = %d\n", LowerBound(sorted, 3))
	fmt.Printf("LowerBound(_, 4)              = %d\n", LowerBound(sorted, 4))
	fmt.Printf("UpperBound(_, 3)              = %d\n", UpperBound(sorted, 3))
	fmt.Printf("UpperBound(_, 7)              = %d\n", UpperBound(sorted, 7))

	fmt.Printf("SearchRange(_, 7)             = %v\n", SearchRange(sorted, 7))
	fmt.Printf("SearchRange(_, 4)             = %v\n", SearchRange(sorted, 4))
	fmt.Printf("CountOccurrences(_, 7)        = %d\n", CountOccurrences(sorted, 7))
	fmt.Printf("CountOccurrences(_, 4)        = %d\n", CountOccurrences(sorted, 4))
}
