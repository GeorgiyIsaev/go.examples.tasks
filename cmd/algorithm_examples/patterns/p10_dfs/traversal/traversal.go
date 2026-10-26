package main

// Паттерн: DFS по дереву — обход в глубину.
//
// Дерево — идеальная структура для рекурсии: у каждого узла ровно
// поддеревья, базис рекурсии — пустое дерево (nil).
//
// Три классических порядка (pre / in / post):
//   - PreOrder  (корень -> лево -> право):  копирование, сериализация;
//   - InOrder   (лево -> корень -> право):  в BST даёт отсортированный порядок;
//   - PostOrder (лево -> право -> корень):  удаление, вычисления снизу вверх
//     (высота, сумма, max path sum).
//
// Рекурсия яснее, итерация (явный стек) нужна, когда глубина дерева
// может упереться в лимит стека или спрашивают «без рекурсии».
//
// Сложность всех обходов: O(n) время, O(h) память (h — высота).

// PreOrder — корень, лево, право (рекурсия).
//
// Пример: [1,2,3,4,5,6,7] -> [1 2 4 5 3 6 7].
func PreOrder(root *TreeNode) []int {
	result := []int{}
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val) // корень первый
		walk(node.Left)
		walk(node.Right)
	}
	walk(root)
	return result
}

// InOrder — лево, корень, право (рекурсия).
// В BST последовательность отсортирована.
//
// Пример: [1,2,3,4,5,6,7] -> [4 2 5 1 6 3 7].
func InOrder(root *TreeNode) []int {
	result := []int{}
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		result = append(result, node.Val) // корень в середине
		walk(node.Right)
	}
	walk(root)
	return result
}

// PostOrder — лево, право, корень (рекурсия).
//
// Пример: [1,2,3,4,5,6,7] -> [4 5 2 6 7 3 1].
func PostOrder(root *TreeNode) []int {
	result := []int{}
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		walk(node.Right)
		result = append(result, node.Val) // корень последний
	}
	walk(root)
	return result
}

// PreOrderIterative — то же, но на явном стеке (без рекурсии).
// Стек: кладём правого потомка, затем левого — тогда левый достанется
// первым (LIFO).
func PreOrderIterative(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		// Правого — первым в стек, чтобы левый вышел раньше.
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// InOrderIterative — лево-корень-право на явном стеке.
// Алгоритм: пока есть левый ребёнок — спускаемся и кладём узел в стек;
// затем разворачиваемся вправо.
func InOrderIterative(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur) // заходим глубже по левой ветке
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}
