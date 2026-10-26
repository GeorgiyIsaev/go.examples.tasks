package main

import "fmt"

func main() {
	fmt.Println("=== 10. DFS / Traversal (обход дерева в глубину) ===")
	fmt.Println()

	root := BuildTree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("дерево: [1 2 3 4 5 6 7] (полное, высота 3)")
	fmt.Println("PreOrder(корень-лево-право)  =", PreOrder(root))  // [1 2 4 5 3 6 7]
	fmt.Println("InOrder(лево-корень-право)   =", InOrder(root))   // [4 2 5 1 6 3 7]
	fmt.Println("PostOrder(лево-право-корень) =", PostOrder(root)) // [4 5 2 6 7 3 1]
	fmt.Println("PreOrder итеративно          =", PreOrderIterative(root))
	fmt.Println("InOrder итеративно           =", InOrderIterative(root))
	fmt.Println()

	bst := BuildTree([]int{4, 2, 6, 1, 3, 5, 7})
	fmt.Println("BST [4 2 6 1 3 5 7], InOrder =", InOrder(bst)) // [1 2 3 4 5 6 7]
	fmt.Println()

	fmt.Println("PreOrder(nil)  =", PreOrder(nil))
	fmt.Println("InOrder(nil)   =", InOrder(nil))
	fmt.Println("PostOrder(nil) =", PostOrder(nil))
}
