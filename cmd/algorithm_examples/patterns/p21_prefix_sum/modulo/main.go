package main

import "fmt"

func main() {
	fmt.Println("=== 21. Prefix Sum / Modulo (остатки) ===")
	fmt.Println()

	fmt.Println("SubarraySumDivByK([4,5,0,-2,-3,1], 5) =",
		SubarraySumDivByK([]int{4, 5, 0, -2, -3, 1}, 5)) // 7
	fmt.Println("SubarraySumDivByK([5], 9) =",
		SubarraySumDivByK([]int{5}, 9)) // 0
	fmt.Println()

	fmt.Println("FindMaxLength([0,1]) =", FindMaxLength([]int{0, 1}))                // 2
	fmt.Println("FindMaxLength([0,1,0]) =", FindMaxLength([]int{0, 1, 0}))           // 2
	fmt.Println("FindMaxLength([0,1,1,0,1]) =", FindMaxLength([]int{0, 1, 1, 0, 1})) // 4
	fmt.Println("FindMaxLength([0]) =", FindMaxLength([]int{0}))                     // 0
}
