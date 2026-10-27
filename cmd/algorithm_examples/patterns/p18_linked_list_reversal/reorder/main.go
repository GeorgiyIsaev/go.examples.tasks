package main

import "fmt"

func main() {
	fmt.Println("=== 18. Linked List Reversal / Reorder ===")
	fmt.Println()

	for _, vals := range [][]int{{1, 2, 3, 4}, {1, 2, 3, 4, 5}} {
		head := BuildList(vals)
		ReorderList(head)
		fmt.Printf("ReorderList(%v) = %v\n", vals, ListToSlice(head))
	}

	single := BuildList([]int{1})
	ReorderList(single)
	fmt.Println("ReorderList([1]) =", ListToSlice(single))

	ReorderList(nil)
	fmt.Println("ReorderList(nil) не паникует")
}
