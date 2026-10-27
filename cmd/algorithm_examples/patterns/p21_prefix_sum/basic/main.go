package main

import "fmt"

func main() {
	fmt.Println("=== 21. Prefix Sum / Basic ===")
	fmt.Println()

	fmt.Println("RunningSum([1,2,3,4]) =", RunningSum([]int{1, 2, 3, 4}))           // [1 3 6 10]
	fmt.Println("PivotIndex([1,7,3,6,5,6]) =", PivotIndex([]int{1, 7, 3, 6, 5, 6})) // 3
	fmt.Println("PivotIndex([1,2,3]) =", PivotIndex([]int{1, 2, 3}))                // -1
	fmt.Println()

	q := NewRangeSumQuery([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println("RangeSumQuery [-2 0 3 -5 2 -1]")
	fmt.Println("SumRange(0,2) =", q.SumRange(0, 2)) // 1
	fmt.Println("SumRange(2,5) =", q.SumRange(2, 5)) // -1
	fmt.Println("SumRange(0,5) =", q.SumRange(0, 5)) // -3
}
