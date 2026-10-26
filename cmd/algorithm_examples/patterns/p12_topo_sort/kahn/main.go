package main

import "fmt"

func main() {
	fmt.Println("=== 12. Topological Sort / Kahn (очередь по indegree) ===")
	fmt.Println()

	// 0 -> 1 -> 3, 0 -> 2 -> 3
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	fmt.Println("порядок [0,1,3],[0,2,3]:", TopoSortKahn(4, edges))
	fmt.Println("длиной:", len(TopoSortKahn(4, edges)))
	fmt.Println()

	cycle := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println("с циклом 0->1->2->0:", TopoSortKahn(3, cycle), "(цикл!)")
	fmt.Println()

	fmt.Println("CanFinish(2, [[1,0]]) =", CanFinish(2, [][]int{{1, 0}}))               // true
	fmt.Println("CanFinish(2, [[1,0],[0,1]]) =", CanFinish(2, [][]int{{1, 0}, {0, 1}})) // false
	fmt.Println()

	fmt.Println("FindOrder(4, [[1,0],[2,0],[3,1],[3,2]]) =",
		FindOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
