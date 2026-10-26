package main

// Паттерн: BFS по уровням на бинарном дереве.
//
// Идея: очередь. Кладём корень, достаём узел, добавляем его детей.
// Чтобы разбить на УРОВНИ, перед обработкой запоминаем длину очереди —
// ровно столько узлов и составляет текущий уровень.
//
// Когда применять:
//   - «обход по уровням», «уровень дерева»;
//   - «минимальная глубина», «самый близкий лист»;
//   - «среднее по уровню», «сумма уровня»;
//   - в ширину — первый увиденный узел = ближайший.
//
// Сложность: O(n) время, O(w) память, где w — максимальная ширина дерева.

// LevelOrder возвращает значения узлов, сгруппированные по уровням.
//
// Пример: дерево 3/9/20/15/7 -> [[3],[9,20],[15,7]].
func LevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue) // узлов текущего уровня
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}
	return result
}

// AverageOfLevels — среднее значение узлов на каждом уровне.
//
// Пример: 1/2/3 -> [1, 2.5].
func AverageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

// MinDepth — минимальная глубина дерева (расстояние от корня до
// самого ближнего листа). Именно для минимальной глубины BFS подходит
// лучше рекурсии: первый достигнутый лист и есть самый близкий.
//
// Пример: 3/9/20/15/7 -> 2. Пустое дерево -> 0.
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Лист — у него нет детей.
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}

// MaxLevelSum — сумма значений каждого уровня; возвращает номер уровня
// (с 1) с максимальной суммой. При равенстве — меньший номер.
//
// Пример: 1/7/0/7/-8 -> уровень 2 (сумма 7).
func MaxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	bestLevel, bestSum := 0, 0
	level := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level++
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level == 1 || sum > bestSum {
			bestSum, bestLevel = sum, level
		}
	}
	return bestLevel
}
