package main

import "fmt"

// Запуск: go run ./tree_level
func main() {
	fmt.Println("=== 09. BFS / Tree Level (обход дерева по уровням) ===")
	fmt.Println()

	root := BuildTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Printf("дерево: %v\n", treeSlice(root))
	fmt.Printf("LevelOrder(root)       = %v\n", LevelOrder(root))
	fmt.Printf("AverageOfLevels(root)  = %v\n", AverageOfLevels(root))
	fmt.Printf("MinDepth(root)         = %d\n", MinDepth(root))
	fmt.Printf("MaxLevelSum(root)      = %d\n", MaxLevelSum(root))

	fmt.Println()

	skewed := BuildTree([]int{1, 2, 3, 4, -1, -1, 5})
	fmt.Printf("дерево: %v\n", treeSlice(skewed))
	fmt.Printf("LevelOrder(skewed)     = %v\n", LevelOrder(skewed))
	fmt.Printf("MinDepth(skewed)       = %d\n", MinDepth(skewed))

	fmt.Println()

	fmt.Printf("LevelOrder(nil)        = %v\n", LevelOrder(nil))
	fmt.Printf("MinDepth(nil)          = %d\n", MinDepth(nil))
}
