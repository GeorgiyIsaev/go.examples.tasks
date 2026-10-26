package main

import "fmt"

func main() {
	graph := BuildGraph([][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}, {3, 4}})
	fmt.Println("Обход с 0:", BFS(graph, 0))                   // [0 1 2 3 4]
	fmt.Println("Путь 0->4:", ShortestPathLength(graph, 0, 4)) // 3

	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	fmt.Println("Есть путь 0->3:", HasPath(edges, 0, 3)) // true
	fmt.Println("Есть путь 0->9:", HasPath(edges, 0, 9)) // false

	fmt.Println("Во все комнаты:", CanVisitAllRooms([][]int{{1}, {2}, {3}, {}}))      // true
	fmt.Println("Не во все:", CanVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}})) // false
}
