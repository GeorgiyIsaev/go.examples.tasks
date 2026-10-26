package main

import "fmt"

func main() {
	fmt.Println("=== 16. DP 2D / Grid (пути по сетке) ===")
	fmt.Println()

	fmt.Println("UniquePaths(3,7) =", UniquePaths(3, 7)) // 28
	fmt.Println("UniquePaths(3,2) =", UniquePaths(3, 2)) // 3
	fmt.Println("UniquePaths(1,1) =", UniquePaths(1, 1)) // 1
	fmt.Println()

	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println("MinPathSum =", MinPathSum(grid)) // 7
	fmt.Println("MinPathSum 1x1 =", MinPathSum([][]int{{5}}))
}
