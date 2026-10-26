package main

import "fmt"

func main() {
	fmt.Println("=== 10. DFS / Graph DFS (обход графа в глубину) ===")
	fmt.Println()

	line := BuildGraph([][]int{{0, 1}, {1, 2}, {2, 3}})
	fmt.Println("DFS с 0 по линии 0-1-2-3 :", DFS(line, 0))
	fmt.Println("DFS итеративно           :", DFSIterative(line, 0))

	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println("Есть путь 0->3 (DFS)     :", HasPathDFS(edges, 0, 3)) // true
	fmt.Println("Есть путь 0->9 (DFS)     :", HasPathDFS(edges, 0, 9)) // false
	fmt.Println()

	fmt.Println("Цикл в квадрате 0-1-2-3-0 :", HasCycle([][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}))         // true
	fmt.Println("Цикл в линии 0-1-2        :", HasCycle([][]int{{0, 1}, {1, 2}}))                         // false
	fmt.Println("Цикл в двух компонентах   :", HasCycle([][]int{{0, 1}, {1, 2}, {3, 4}, {4, 5}, {5, 3}})) // true
	fmt.Println()

	fmt.Println("Компонент [[0,1],[2,3]] :", CountComponents(4, [][]int{{0, 1}, {2, 3}})) // 2
	fmt.Println("Компонент [[0,1],[1,2]] :", CountComponents(4, [][]int{{0, 1}, {1, 2}})) // 2
	fmt.Println("Компонент без рёбер n=3 :", CountComponents(3, nil))                     // 3
}
