package main

import "fmt"

func main() {
	fmt.Println("=== 21. Prefix Sum / Products ===")
	fmt.Println()

	fmt.Println("ProductExceptSelf([1,2,3,4]) =", ProductExceptSelf([]int{1, 2, 3, 4})) // [24 12 8 6]
	fmt.Println("ProductExceptSelf([-1,1,0,-3,3]) =", ProductExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println("ProductExceptSelf([2]) =", ProductExceptSelf([]int{2}))
}
