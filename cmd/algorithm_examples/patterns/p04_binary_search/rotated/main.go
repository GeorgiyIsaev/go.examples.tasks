package main

import "fmt"

// Запуск: go run ./rotated
func main() {
	fmt.Println("=== 04. Binary Search / Rotated (повёрнутый массив) ===")
	fmt.Println()

	rotated := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("массив: %v\n", rotated)
	fmt.Printf("SearchRotated(_, 0)           = %d\n", SearchRotated(rotated, 0))
	fmt.Printf("SearchRotated(_, 5)           = %d\n", SearchRotated(rotated, 5))
	fmt.Printf("SearchRotated(_, 3)           = %d\n", SearchRotated(rotated, 3))
	fmt.Printf("FindMinRotated(_)             = %d\n", FindMinRotated(rotated))

	fmt.Println()

	notRotated := []int{1, 2, 3, 4, 5}
	fmt.Printf("массив: %v\n", notRotated)
	fmt.Printf("SearchRotated(_, 3)           = %d\n", SearchRotated(notRotated, 3))
	fmt.Printf("FindMinRotated(_)             = %d\n", FindMinRotated(notRotated))
}
