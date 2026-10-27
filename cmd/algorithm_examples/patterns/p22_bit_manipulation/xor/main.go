package main

import "fmt"

func main() {
	fmt.Println("=== 22. Bit Manipulation / XOR ===")
	fmt.Println()

	fmt.Println("SingleNumber([4,1,2,1,2]) =", SingleNumber([]int{4, 1, 2, 1, 2})) // 4
	fmt.Println("SingleNumber([1]) =", SingleNumber([]int{1}))                     // 1
	fmt.Println()

	fmt.Println("SingleNumberII([2,2,3,2]) =", SingleNumberII([]int{2, 2, 3, 2}))                  // 3
	fmt.Println("SingleNumberII([0,1,0,1,0,1,99]) =", SingleNumberII([]int{0, 1, 0, 1, 0, 1, 99})) // 99
	fmt.Println()

	fmt.Println("SingleNumberIII([1,2,1,3,2,5]) =", SingleNumberIII([]int{1, 2, 1, 3, 2, 5})) // [3 5]
	fmt.Println("SingleNumberIII([-1,0]) =", SingleNumberIII([]int{-1, 0}))                   // [-1 0]
}
