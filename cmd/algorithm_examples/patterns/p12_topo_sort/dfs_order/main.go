package main

import "fmt"

func main() {
	fmt.Println("=== 12. Topological Sort / DFS (обратный порядок окончания) ===")
	fmt.Println()

	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}
	fmt.Println("TopoSortDFS [0->1,0->2,1->3,2->3]:", TopoSortDFS(4, edges))
	fmt.Println()

	cycle := [][]int{{0, 1}, {1, 2}, {2, 0}}
	fmt.Println("с циклом:", TopoSortDFS(3, cycle), "(цикл!)")
	fmt.Println()

	fmt.Println("CanFinishDFS(2, [[1,0]])        =", CanFinishDFS(2, [][]int{{1, 0}}))         // true
	fmt.Println("CanFinishDFS(2, [[1,0],[0,1]])  =", CanFinishDFS(2, [][]int{{1, 0}, {0, 1}})) // false
	fmt.Println()

	fmt.Println("HasCycleDirected(4, [[0,1],[1,2],[2,3]]) =", HasCycleDirected(4, [][]int{{0, 1}, {1, 2}, {2, 3}})) // false
	fmt.Println("HasCycleDirected(3, [[0,1],[1,2],[2,0]]) =", HasCycleDirected(3, [][]int{{0, 1}, {1, 2}, {2, 0}})) // true
}
