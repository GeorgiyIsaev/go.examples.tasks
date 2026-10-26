package main

import "fmt"

func main() {
	fmt.Println("=== 12. Topological Sort / Queries (предшественник?) ===")
	fmt.Println()

	// [a, b] — «a идёт до b»: 1->0, 2->0, 3->1, 3->2
	prereq := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	queries := [][]int{{0, 1}, {0, 3}, {3, 0}, {3, 1}}
	// ожидаем: false, false, true, true

	fmt.Println("граф: 3->1->0, 3->2->0")
	fmt.Println("запросы [0,1],[0,3],[3,0],[3,1]:")
	fmt.Println("  DFS на каждый запрос :", CheckIfPrerequisite(4, prereq, queries))
	fmt.Println("  таблица достижимости :", PredecessorQueries(4, prereq, queries))
}
