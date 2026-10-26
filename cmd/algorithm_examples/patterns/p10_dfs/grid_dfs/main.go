package main

import "fmt"

func main() {
	fmt.Println("=== 10. DFS / Grid DFS (заливка по сетке) ===")
	fmt.Println()

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println("Островов (DFS):", NumIslandsDFS(grid)) // 3

	island := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println("Площадь самого большого:", MaxAreaOfIsland(island)) // 6

	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	filled := FloodFill(image, 1, 1, 2)
	fmt.Println("Заливка (1,1) -> 2:")
	for _, row := range filled {
		fmt.Println(" ", row)
	}
}
