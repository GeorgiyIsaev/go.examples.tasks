package main

import "fmt"

// Запуск: go run ./on_answer
func main() {
	fmt.Println("=== 04. Binary Search / On Answer (по ответу) ===")
	fmt.Println()

	fmt.Printf("MySqrt(4)                     = %d\n", MySqrt(4))
	fmt.Printf("MySqrt(8)                     = %d\n", MySqrt(8))
	fmt.Printf("MySqrt(0)                     = %d\n", MySqrt(0))
	fmt.Printf("MySqrt(100)                   = %d\n", MySqrt(100))

	fmt.Printf("MinEatingSpeed([3,6,7,11], 8)         = %d\n",
		MinEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Printf("MinEatingSpeed([30,11,23,4,20], 5)    = %d\n",
		MinEatingSpeed([]int{30, 11, 23, 4, 20}, 5))

	fmt.Printf("ShipWithinDays([1..10], 5)            = %d\n",
		ShipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Printf("ShipWithinDays([3,2,2,4,1,4], 3)      = %d\n",
		ShipWithinDays([]int{3, 2, 2, 4, 1, 4}, 3))
}
