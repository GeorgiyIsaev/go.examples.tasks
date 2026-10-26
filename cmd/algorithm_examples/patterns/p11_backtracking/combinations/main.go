package main

import "fmt"

func main() {
	fmt.Println("=== 11. Backtracking / Combinations (комбинации и суммы) ===")
	fmt.Println()

	fmt.Println("Combine(n=4, k=2):")
	for _, c := range Combine(4, 2) {
		fmt.Println(" ", c)
	}
	fmt.Println("всего:", len(Combine(4, 2)), "(C(4,2)=6)")
	fmt.Println()

	fmt.Println("CombinationSum([2,3,6,7], 7):")
	for _, c := range CombinationSum([]int{2, 3, 6, 7}, 7) {
		fmt.Println(" ", c)
	}
	fmt.Println()

	fmt.Println("CombinationSumII([10,1,2,7,6,1,5], 8):")
	for _, c := range CombinationSumII([]int{10, 1, 2, 7, 6, 1, 5}, 8) {
		fmt.Println(" ", c)
	}
}
