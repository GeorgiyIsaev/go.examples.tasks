package main

import "fmt"

func main() {
	fmt.Println("=== 05. Hash Map / Pairs (пара с суммой) ===")
	fmt.Println()

	nums := []int{2, 7, 11, 15}
	fmt.Println("TwoSum([2,7,11,15], 9) =", TwoSum(nums, 9)) // [0 1]
	fmt.Println("TwoSum([1,2,3], 100)   =", TwoSum(nums, 100), "(нет пары)")
	fmt.Println()

	fmt.Println("TwoSumCount([1,1,1], 2) =", TwoSumCount([]int{1, 1, 1}, 2)) // 3
	fmt.Println("TwoSumCount([1,2,3], 3) =", TwoSumCount([]int{1, 2, 3}, 3)) // 1
	fmt.Println()

	fmt.Println("SubarraySum([1,1,1], 2)   =", SubarraySum([]int{1, 1, 1}, 2))  // 2
	fmt.Println("SubarraySum([1,2,3], 3)   =", SubarraySum([]int{1, 2, 3}, 3))  // 2
	fmt.Println("SubarraySum([1,-1,0], 0)  =", SubarraySum([]int{1, -1, 0}, 0)) // 3
}
