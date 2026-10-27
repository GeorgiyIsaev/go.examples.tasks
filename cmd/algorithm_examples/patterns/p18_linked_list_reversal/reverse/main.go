package main

import "fmt"

func main() {
	fmt.Println("=== 18. Linked List Reversal / Reverse ===")
	fmt.Println()

	head := BuildList([]int{1, 2, 3, 4, 5})
	fmt.Println("исходный:      ", ListToSlice(head))
	fmt.Println("итеративно:    ", ListToSlice(ReverseList(BuildList([]int{1, 2, 3, 4, 5}))))
	fmt.Println("рекурсивно:    ", ListToSlice(ReverseListRecursive(BuildList([]int{1, 2, 3, 4, 5}))))
	fmt.Println("пустой:        ", ReverseList(nil), "(nil)")
	fmt.Println()

	fmt.Println("ReverseBetween(1..5, 2, 4):",
		ListToSlice(ReverseBetween(BuildList([]int{1, 2, 3, 4, 5}), 2, 4))) // [1 4 3 2 5]
	fmt.Println("ReverseBetween(1..5, 1, 5):",
		ListToSlice(ReverseBetween(BuildList([]int{1, 2, 3, 4, 5}), 1, 5))) // [5 4 3 2 1]
	fmt.Println("ReverseBetween(1..5, 3, 3):",
		ListToSlice(ReverseBetween(BuildList([]int{1, 2, 3, 4, 5}), 3, 3))) // без изменений
}
