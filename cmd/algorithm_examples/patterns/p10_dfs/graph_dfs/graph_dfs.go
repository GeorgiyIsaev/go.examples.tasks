package main

// Паттерн: DFS по графу (рекурсия или явный стек).
//
// Идея: от стартовой вершины уходим как можно глубже по одному
// соседу, помечая посещённые; вернувшись, пробуем следующего.
// Отличие от BFS: вместо очереди — стек (и «в глубину», а не слоями).
//
// Граф здесь — список смежности map[int][]int (как в p09_bfs/graph_bfs).
// Сложность: O(V + E).
//
// Когда DFS предпочтительнее BFS:
//   - проверка связности / существования пути (early exit, без уровней);
//   - поиск циклов;
//   - топологическая сортировка;
//   - задачи, где нужен путь/подмножество, а не кратчайшее расстояние.

// BuildGraph строит неориентированный список смежности из рёбер.
func BuildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

// DFS — обход в глубину от start, возвращает порядок посещения.
// Рекурсивная реализация.
//
// Пример: линия 0-1-2-3, старт 0 -> [0 1 2 3].
func DFS(graph map[int][]int, start int) []int {
	if _, ok := graph[start]; !ok {
		return []int{}
	}
	visited := map[int]bool{}
	order := []int{}

	var visit func(node int)
	visit = func(node int) {
		visited[node] = true
		order = append(order, node)
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visit(neighbor)
			}
		}
	}
	visit(start)
	return order
}

// DFSIterative — то же на явном стеке (без рекурсии; глубина не
// ограничена размером стека вызовов).
// Порядок соседей получается обратным из-за LIFO — для порядка
// посещения это не критично, тесты проверяют множество узлов.
func DFSIterative(graph map[int][]int, start int) []int {
	if _, ok := graph[start]; !ok {
		return []int{}
	}
	visited := map[int]bool{start: true}
	stack := []int{start}
	order := []int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		order = append(order, node)

		// Обратный порядок, чтобы соседи брались слева направо.
		for i := len(graph[node]) - 1; i >= 0; i-- {
			neighbor := graph[node][i]
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			stack = append(stack, neighbor)
		}
	}
	return order
}

// HasPathDFS — существует ли путь между start и end (LeetCode 1971),
// обход в глубину с ранним выходом.
func HasPathDFS(edges [][]int, start, end int) bool {
	graph := BuildGraph(edges)
	if start == end {
		return true
	}
	if _, ok := graph[start]; !ok {
		return false
	}
	visited := map[int]bool{start: true}

	var search func(node int) bool
	search = func(node int) bool {
		for _, neighbor := range graph[node] {
			if neighbor == end {
				return true
			}
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			if search(neighbor) {
				return true
			}
		}
		return false
	}
	return search(start)
}

// HasCycle — есть ли цикл в неориентированном графе (UF-классика,
// но здесь DFS-вариант).
//
// Трюк: помнить «родителя» — если сосед уже посещён и это не родитель,
// значит, ребро ведёт назад => цикл.
//
// Пример: квадрат 0-1-2-3-0 -> true; линия 0-1-2 -> false.
func HasCycle(edges [][]int) bool {
	graph := BuildGraph(edges)
	visited := map[int]bool{}

	var walk func(node, parent int) bool
	walk = func(node, parent int) bool {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if neighbor == parent {
				continue // ребро, по которому пришли
			}
			if visited[neighbor] {
				return true // встретили посещённого не-родителя
			}
			if walk(neighbor, node) {
				return true
			}
		}
		return false
	}

	// Граф может быть несвязным — обходим все компоненты.
	for node := range graph {
		if !visited[node] {
			if walk(node, -1) {
				return true
			}
		}
	}
	return false
}

// CountComponents — количество связных компонент графа
// (LeetCode 323 / 547).
//
// Пример: рёбра [[0,1],[2,3]] -> 2 компоненты {0,1} и {2,3}.
func CountComponents(n int, edges [][]int) int {
	graph := BuildGraph(edges)
	visited := map[int]bool{}
	count := 0

	for node := 0; node < n; node++ {
		if visited[node] || len(graph[node]) == 0 {
			// изолированная вершина — своя компонента, но может не
			// попадать в граф
			if !visited[node] {
				visited[node] = true
				count++
			}
			continue
		}
		// DFS по компоненте.
		stack := []int{node}
		visited[node] = true
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, neighbor := range graph[cur] {
				if !visited[neighbor] {
					visited[neighbor] = true
					stack = append(stack, neighbor)
				}
			}
		}
		count++
	}
	return count
}
