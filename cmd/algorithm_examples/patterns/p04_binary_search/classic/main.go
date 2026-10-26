package main

import "fmt"

// Запуск: go run ./classic
func main() {
	fmt.Println("=== 04. Binary Search / Classic ===")
	fmt.Println()

	nums := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Printf("BinarySearch(%v, 7)           = %d\n", nums, BinarySearch(nums, 7))
	fmt.Printf("BinarySearch(%v, 4)           = %d\n", nums, BinarySearch(nums, 4))

	ins := []int{1, 3, 5, 6}
	fmt.Printf("SearchInsert(%v, 5)           = %d\n", ins, SearchInsert(ins, 5))
	fmt.Printf("SearchInsert(%v, 2)           = %d\n", ins, SearchInsert(ins, 2))
	fmt.Printf("SearchInsert(%v, 7)           = %d\n", ins, SearchInsert(ins, 7))
	fmt.Printf("SearchInsert(%v, 0)           = %d\n", ins, SearchInsert(ins, 0))

	// Симулируем FirstBadVersion: firstBad = 4, n = 5
	firstBad := 4
	isBad := func(i int) bool { return i >= firstBad }
	fmt.Printf("FirstBadVersion(n=5, firstBad=%d) = %d\n", firstBad, FirstBadVersion(5, isBad))
}
