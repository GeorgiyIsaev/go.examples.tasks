package main

import "fmt"

// Запуск: go run ./happy_number
func main() {
	fmt.Println("=== 02. Fast / Slow / Happy Number (Floyd's на числах) ===")
	fmt.Println()

	for _, n := range []int{19, 1, 7, 10, 100, 2, 3, 4} {
		fmt.Printf("IsHappy(%3d) = %v\n", n, IsHappy(n))
	}
}
