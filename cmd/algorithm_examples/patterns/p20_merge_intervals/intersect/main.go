package main

import "fmt"

func main() {
	fmt.Println("=== 20. Merge Intervals / Intersection ===")
	fmt.Println()

	a := [][]int{{0, 2}, {5, 10}, {13, 15}, {16, 18}}
	b := [][]int{{1, 5}, {8, 12}, {15, 24}}
	fmt.Println("A =", a)
	fmt.Println("B =", b)
	fmt.Println("A ∩ B =", IntervalIntersection(a, b))
	fmt.Println()

	fmt.Println("без пересечений =", IntervalIntersection(
		[][]int{{1, 3}}, [][]int{{5, 7}})) // []
	fmt.Println("касание =", IntervalIntersection(
		[][]int{{1, 5}}, [][]int{{5, 8}})) // [[5 5]]
}
