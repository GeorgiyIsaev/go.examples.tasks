package main

import "fmt"

func main() {
	fmt.Println("=== 20. Merge Intervals / Summary Ranges ===")
	fmt.Println()

	fmt.Println("SummaryRanges([0,1,2,4,5,7]) =",
		SummaryRanges([]int{0, 1, 2, 4, 5, 7})) // [0->2 4->5 7]
	fmt.Println("SummaryRanges([0,2,3,4,6,8,9]) =",
		SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // [0 2->4 6 8->9]
	fmt.Println("SummaryRanges([]) =", SummaryRanges([]int{}))
	fmt.Println("SummaryRanges([-1]) =", SummaryRanges([]int{-1}))
}
