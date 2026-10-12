package main

import "fmt"

func main() {
	fmt.Println("=== 04. Binary Search ===")
	fmt.Println()

	fmt.Println("--- Classic ---")
	nums := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Printf("BinarySearch(%v, 7)                 = %d\n", nums, BinarySearch(nums, 7))
	fmt.Printf("BinarySearch(%v, 4)                 = %d\n", nums, BinarySearch(nums, 4))
	fmt.Printf("SearchInsert([1,3,5,6], 5)         = %d\n", SearchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Printf("SearchInsert([1,3,5,6], 2)         = %d\n", SearchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Printf("SearchInsert([1,3,5,6], 7)         = %d\n", SearchInsert([]int{1, 3, 5, 6}, 7))

	// Симулируем FirstBadVersion: firstBad = 4, n = 5
	firstBad := 4
	isBad := func(i int) bool { return i >= firstBad }
	fmt.Printf("FirstBadVersion(5, isBad=%d)        = %d\n", firstBad, FirstBadVersion(5, isBad))

	fmt.Println()
	fmt.Println("--- Bounds ---")
	sorted := []int{1, 3, 3, 5, 7, 7, 7, 9}
	fmt.Printf("LowerBound(%v, 3)                   = %d\n", sorted, LowerBound(sorted, 3))
	fmt.Printf("UpperBound(%v, 3)                   = %d\n", sorted, UpperBound(sorted, 3))
	fmt.Printf("SearchRange(%v, 7)                  = %v\n", sorted, SearchRange(sorted, 7))
	fmt.Printf("SearchRange(%v, 4)                  = %v\n", sorted, SearchRange(sorted, 4))
	fmt.Printf("CountOccurrences(%v, 7)             = %d\n", sorted, CountOccurrences(sorted, 7))

	fmt.Println()
	fmt.Println("--- Rotated ---")
	rotated := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("SearchRotated(%v, 0)                = %d\n", rotated, SearchRotated(rotated, 0))
	fmt.Printf("SearchRotated(%v, 5)                = %d\n", rotated, SearchRotated(rotated, 5))
	fmt.Printf("SearchRotated(%v, 3)                = %d\n", rotated, SearchRotated(rotated, 3))
	fmt.Printf("FindMinRotated(%v)                  = %d\n", rotated, FindMinRotated(rotated))

	fmt.Println()
	fmt.Println("--- Binary Search on Answer ---")
	fmt.Printf("MySqrt(4)                            = %d\n", MySqrt(4))
	fmt.Printf("MySqrt(8)                            = %d\n", MySqrt(8))
	fmt.Printf("MySqrt(0)                            = %d\n", MySqrt(0))
	fmt.Printf("MinEatingSpeed([3,6,7,11], 8)        = %d\n",
		MinEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Printf("MinEatingSpeed([30,11,23,4,20], 5)   = %d\n",
		MinEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Printf("ShipWithinDays([1..10], 5)           = %d\n",
		ShipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
