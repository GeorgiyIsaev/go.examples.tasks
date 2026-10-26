package main

import "fmt"

func main() {
	fmt.Println("=== 08. Monotonic Stack / Next Greater ===")
	fmt.Println()

	fmt.Println("nextGreater([2,1,2,4,3]) =", nextGreater([]int{2, 1, 2, 4, 3})) // [4 2 4 -1 -1]
	fmt.Println()

	fmt.Println("NextGreaterElement([4,1,2],[1,3,4,2]) =",
		NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})) // [-1 3 -1]
	fmt.Println()

	fmt.Println("NextGreaterElements([1,2,1]) =", NextGreaterElements([]int{1, 2, 1})) // [2 -1 2]
	fmt.Println("NextGreaterElements([4,3,2,1]) =", NextGreaterElements([]int{4, 3, 2, 1}))
}
