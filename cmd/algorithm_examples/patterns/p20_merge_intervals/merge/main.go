package main

import "fmt"

func main() {
	fmt.Println("=== 20. Merge Intervals / Merge & Insert ===")
	fmt.Println()

	fmt.Println("Merge([[1,3],[2,6],[8,10],[15,18]]) =",
		Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println("Merge([[1,4],[4,5]]) =",
		Merge([][]int{{1, 4}, {4, 5}})) // [[1 5]] — касание слияется
	fmt.Println()

	fmt.Println("Insert([[1,3],[6,9]], [2,5]) =",
		Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})) // [[1 5] [6 9]]
	fmt.Println("Insert([[1,3],[6,9]], [10,12]) =",
		Insert([][]int{{1, 3}, {6, 9}}, []int{10, 12})) // [[1 3] [6 9] [10 12]]
	fmt.Println("Insert([], [5,7]) =", Insert(nil, []int{5, 7})) // [[5 7]]
}
