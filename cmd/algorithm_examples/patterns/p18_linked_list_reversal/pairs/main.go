package main

import "fmt"

func main() {
	fmt.Println("=== 18. Linked List Reversal / Pairs & K-Groups ===")
	fmt.Println()

	fmt.Println("SwapPairs(1..5):",
		ListToSlice(SwapPairs(BuildList([]int{1, 2, 3, 4, 5})))) // [2 1 4 3 5]
	fmt.Println("SwapPairs(1..4):",
		ListToSlice(SwapPairs(BuildList([]int{1, 2, 3, 4})))) // [2 1 4 3]
	fmt.Println("SwapPairs(1):", ListToSlice(SwapPairs(BuildList([]int{1}))))
	fmt.Println()

	fmt.Println("ReverseKGroup(1..5, k=2):",
		ListToSlice(ReverseKGroup(BuildList([]int{1, 2, 3, 4, 5}), 2))) // [2 1 4 3 5]
	fmt.Println("ReverseKGroup(1..5, k=3):",
		ListToSlice(ReverseKGroup(BuildList([]int{1, 2, 3, 4, 5}), 3))) // [3 2 1 4 5]
	fmt.Println("ReverseKGroup(1..4, k=4):",
		ListToSlice(ReverseKGroup(BuildList([]int{1, 2, 3, 4}), 4))) // [4 3 2 1]
	fmt.Println("ReverseKGroup(1..3, k=5):",
		ListToSlice(ReverseKGroup(BuildList([]int{1, 2, 3}), 5))) // [1 2 3] — хвост не трогаем
}
