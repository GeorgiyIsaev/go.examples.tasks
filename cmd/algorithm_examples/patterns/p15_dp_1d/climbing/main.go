package main

import "fmt"

func main() {
	fmt.Println("=== 15. DP 1D / Climbing Stairs ===")
	fmt.Println()

	for n := 1; n <= 6; n++ {
		fmt.Printf("ClimbStairs(%d) = %d\n", n, ClimbStairs(n))
	}
	fmt.Println()

	for _, c := range [][]int{{10, 15, 20}, {1, 100, 1, 1, 1, 100, 1, 1, 100, 1}} {
		fmt.Printf("MinCostClimbingStairs(%v) = %d\n", c, MinCostClimbingStairs(c))
	}
}
