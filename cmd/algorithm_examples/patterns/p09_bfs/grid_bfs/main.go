package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println("Островов:", NumIslands(grid)) // 3

	rotting := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println("Минут до гниения:", RottingOranges(rotting)) // 4

	unreachable := [][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	}
	fmt.Println("С непроходимым:", RottingOranges(unreachable)) // -1

	path := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}
	fmt.Println("Кратчайший путь:", ShortestPathBinaryMatrix(path)) // 4
}
