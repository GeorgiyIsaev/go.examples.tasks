package main

import "fmt"

func main() {
	fmt.Println("=== 17. Heap / Merge K (слияние k списков) ===")
	fmt.Println()

	lists := []*ListNode{
		BuildList([]int{1, 4, 5}),
		BuildList([]int{1, 3, 4}),
		BuildList([]int{2, 6}),
	}
	merged := MergeKLists(lists)
	fmt.Println("MergeKLists([1,4,5],[1,3,4],[2,6]) =", ListToSlice(merged)) // [1 1 2 3 4 4 5 6]
	fmt.Println()

	fmt.Println("с пустым списком среди:", ListToSlice(MergeKLists([]*ListNode{
		BuildList([]int{1, 2}), nil, BuildList([]int{3}),
	}))) // [1 2 3]
	fmt.Println("все пустые:", MergeKLists([]*ListNode{nil, nil}), "(nil)")
}
