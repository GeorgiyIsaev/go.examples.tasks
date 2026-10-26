package main

import "fmt"

// Запуск: go run ./fixed_window
func main() {
	fmt.Println("=== 02. Sliding Window / Fixed (фиксированное окно) ===")
	fmt.Println()

	fmt.Printf("MaxSumSubarrayOfSizeK([2,1,5,1,3,2], 3)   = %d\n",
		MaxSumSubarrayOfSizeK([]int{2, 1, 5, 1, 3, 2}, 3))
	fmt.Printf("MaxSumSubarrayOfSizeK([1,2], 5)            = %d\n",
		MaxSumSubarrayOfSizeK([]int{1, 2}, 5))

	fmt.Printf("FindAverages([1,3,2,6,-1,4,1,8,2], 5)     = %v\n",
		FindAverages([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))

	fmt.Printf("MaxVowels(\"abciiidef\", 3)                  = %d\n",
		MaxVowels("abciiidef", 3))
	fmt.Printf("MaxVowels(\"bcdfg\", 2)                      = %d\n",
		MaxVowels("bcdfg", 2))
}
