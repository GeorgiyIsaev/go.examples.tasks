package main

import "fmt"

// Запуск: go run ./three_pointers
func main() {
	fmt.Println("=== 01. Two Pointers / Three Pointers ===")
	fmt.Println()

	fmt.Printf("ThreeSum([-1,0,1,2,-1,-4]) = %v\n", ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("ThreeSum([0,0,0])          = %v\n", ThreeSum([]int{0, 0, 0}))
	fmt.Printf("ThreeSum([1,2,3])          = %v\n", ThreeSum([]int{1, 2, 3}))
	fmt.Printf("ThreeSum([-2,0,0,2,2])     = %v\n", ThreeSum([]int{-2, 0, 0, 2, 2}))

	fmt.Println()

	colors := []int{2, 0, 2, 1, 1, 0}
	SortColors(colors)
	fmt.Printf("SortColors([2,0,2,1,1,0])  = %v\n", colors)

	colors2 := []int{2, 2, 1, 0, 1, 0}
	SortColors(colors2)
	fmt.Printf("SortColors([2,2,1,0,1,0])  = %v\n", colors2)
}
