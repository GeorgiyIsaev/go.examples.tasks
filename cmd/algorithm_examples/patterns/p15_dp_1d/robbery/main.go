package main

import "fmt"

func main() {
	fmt.Println("=== 15. DP 1D / House Robber ===")
	fmt.Println()

	for _, nums := range [][]int{{1, 2, 3, 1}, {2, 7, 9, 3, 1}, {}} {
		fmt.Printf("Rob(%v) = %d\n", nums, Rob(nums))
	}
	fmt.Println()

	for _, nums := range [][]int{{2, 3, 2}, {1, 2, 3, 1}, {1, 2, 3}, {5}} {
		fmt.Printf("RobCircle(%v) = %d\n", nums, RobCircle(nums))
	}
}
