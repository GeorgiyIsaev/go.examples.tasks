package main

import "fmt"

func main() {
	fmt.Println("=== 14. Greedy / Jump Game ===")
	fmt.Println()

	for _, nums := range [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{0},
		{1, 0},
	} {
		fmt.Printf("CanJump(%v) = %v\n", nums, CanJump(nums))
	}
	fmt.Println()

	for _, nums := range [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
		{1},
	} {
		fmt.Printf("Jump(%v) = %d\n", nums, Jump(nums))
	}
}
