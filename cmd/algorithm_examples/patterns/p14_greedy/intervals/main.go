package main

import "fmt"

func main() {
	fmt.Println("=== 14. Greedy / Intervals (выбор мероприятий) ===")
	fmt.Println()

	fmt.Println("EraseOverlapIntervals([[1,2],[2,3],[3,4],[1,3]]) =",
		EraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // 1
	fmt.Println("EraseOverlapIntervals([[1,2],[1,2],[1,2]]) =",
		EraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}})) // 2
	fmt.Println("EraseOverlapIntervals([[1,2],[2,3]]) =",
		EraseOverlapIntervals([][]int{{1, 2}, {2, 3}})) // 0 — касание ок
	fmt.Println()

	fmt.Println("FindMinArrowShots([[10,16],[2,8],[1,6],[7,12]]) =",
		FindMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})) // 2
	fmt.Println("FindMinArrowShots([[1,2],[3,4],[5,6],[7,8]]) =",
		FindMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}})) // 4
	fmt.Println("FindMinArrowShots([[1,2],[2,3],[3,4],[4,5]]) =",
		FindMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}})) // 2
}
