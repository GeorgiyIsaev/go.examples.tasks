package main

import "fmt"

func main() {
	fmt.Println("=== 17. Heap / Top-K (куча для k-й позиции) ===")
	fmt.Println()

	fmt.Println("KthLargest([3,2,1,5,6,4], 2) =", KthLargest([]int{3, 2, 1, 5, 6, 4}, 2)) // 5
	fmt.Println("KthLargest([3,2,3,1,2,4,5,5,6], 4) =",
		KthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)) // 4
	fmt.Println()

	s := NewKthLargestStream(3, []int{4, 5, 8, 2})
	fmt.Println("поток: k=3, старт [4 5 8 2]")
	for _, v := range []int{3, 5, 10, 9, 4} {
		fmt.Printf("  Add(%2d) -> %d\n", v, s.Add(v))
	}
	fmt.Println()

	fmt.Println("KClosest([[1,3],[-2,2]], 1) =", KClosest([][]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println("KClosest([[3,3],[5,-1],[-2,4]], 2) =",
		KClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
