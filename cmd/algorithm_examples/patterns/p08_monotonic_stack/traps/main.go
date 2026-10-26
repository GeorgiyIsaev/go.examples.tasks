package main

import "fmt"

func main() {
	fmt.Println("=== 08. Monotonic Stack / Traps (вода и гистограмма) ===")
	fmt.Println()

	h := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println("heights:", h)
	fmt.Println("Trap =", Trap(h))                                      // 6
	fmt.Println("Trap([4,2,0,3,2,5]) =", Trap([]int{4, 2, 0, 3, 2, 5})) // 9
	fmt.Println()

	fmt.Println("LargestRectangleArea([2,1,5,6,2,3]) =",
		LargestRectangleArea([]int{2, 1, 5, 6, 2, 3})) // 10
	fmt.Println("LargestRectangleArea([2,4]) =", LargestRectangleArea([]int{2, 4})) // 4
	fmt.Println("LargestRectangleArea([1])   =", LargestRectangleArea([]int{1}))    // 1
}
