package main

import "fmt"

func main() {
	fmt.Println("=== 16. DP 2D / Knapsack (рюкзак) ===")
	fmt.Println()

	fmt.Println("Knapsack01(W=7, w=[1 3 4 5], v=[1 4 5 7]) =",
		Knapsack01([]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7)) // 9
	fmt.Println("Knapsack01(W=0, ...) =",
		Knapsack01([]int{1, 3}, []int{1, 4}, 0)) // 0
	fmt.Println()

	fmt.Println("SubsetSum([1,5,11,5], 11) =", SubsetSum([]int{1, 5, 11, 5}, 11)) // true
	fmt.Println("SubsetSum([1,2,3], 7)     =", SubsetSum([]int{1, 2, 3}, 7))      // false
	fmt.Println()

	fmt.Println("PartitionEqual([1,5,11,5]) =", PartitionEqual([]int{1, 5, 11, 5})) // true
	fmt.Println("PartitionEqual([1,2,3,5])  =", PartitionEqual([]int{1, 2, 3, 5}))  // false
}
