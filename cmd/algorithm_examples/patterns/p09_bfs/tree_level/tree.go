package main

// TreeNode — узел бинарного дерева (стандартное определение LeetCode).
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree строит бинарное дерево из среза в порядке уровня
// (LeetCode-сериализация). Значение -1 означает отсутствие узла.
//
// Пример: [3,9,20,-1,-1,15,7] даёт
//
//	  3
//	 / \
//	9  20
//	   / \
//	  15  7
func BuildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) {
			if vals[i] != -1 {
				node.Left = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Left)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != -1 {
				node.Right = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Right)
			}
			i++
		}
	}
	return root
}

// treeSlice — сериализация обратно в срез уровня (для наглядного вывода).
func treeSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, -1)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Отрезаем хвост из -1.
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}
	return result
}
