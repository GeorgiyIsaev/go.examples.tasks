package main

import "fmt"

// Запуск: go run ./slow_fast
func main() {
	fmt.Println("=== 01. Two Pointers / Slow-Fast (в одном направлении) ===")
	fmt.Println()

	nums1 := []int{1, 1, 2, 3, 3, 3, 4}
	k := RemoveDuplicatesSorted(nums1)
	fmt.Printf("RemoveDuplicatesSorted([1,1,2,3,3,3,4]) = %d, nums = %v\n", k, nums1[:k])

	nums2 := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums2)
	fmt.Printf("MoveZeroes([0,1,0,3,12])                 = %v\n", nums2)

	nums3 := []int{3, 2, 2, 3}
	k3 := RemoveElement(nums3, 3)
	fmt.Printf("RemoveElement([3,2,2,3], 3)              = %d, nums = %v\n", k3, nums3[:k3])

	nums4 := []int{0, 0, 1}
	k4 := RemoveElement(nums4, 0)
	fmt.Printf("RemoveElement([0,0,1], 0)                = %d, nums = %v\n", k4, nums4[:k4])
}
