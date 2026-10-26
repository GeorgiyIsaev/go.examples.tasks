package main

// Паттерн: топологическая сортировка через DFS.
//
// Идея: обходим граф в глубину; когда вершина «закончена» (все её
// преемники уже в порядке) — кладём её в НАЧАЛО результата.
// Итог: обратный порядок окончания DFS = топопорядок.
//
// Отличие от Кана:
//   - Канн даёт порядок «на лету» и цикл ловится по счётчику;
//   - DFS требует состояний вершин, потому что граф может быть
//     несвязным и содержать рёбра «назад» во время обхода.
//
// Три состояния вершины — главный трюк DFS-варианта:
//
//	0 (white)  — ещё не посещена
//	1 (gray)   — ВХОДИЛИ, но ещё не вышли: лежит в текущем стеке вызовов
//	2 (black)  — полностью обработана
//
// Ребро в вершину состояния 1 = мы полезли в самого себя по цепочке
// зависимостей => **цикл**. Ребро в вершину 2 — просто уже готовое
// поддерево, не ошибка.
//
// Сложность: O(V + E), как у Кана.

// TopoSortDFS — топопорядок вершин 0..n-1 через DFS.
// Если цикл — пустой срез.
//
// Пример: рёбра [[0,1],[0,2],[1,3],[2,3]] -> [0 1 2 3] или [0 2 1 3].
func TopoSortDFS(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
	}

	const (
		white = 0 // не посещена
		gray  = 1 // в стеке вызовов
		black = 2 // обработана
	)
	state := make([]int, n) // все white по умолчанию
	order := make([]int, 0, n)
	hasCycle := false

	var visit func(v int)
	visit = func(v int) {
		if hasCycle {
			return
		}
		state[v] = gray
		for _, nb := range graph[v] {
			switch state[nb] {
			case gray:
				hasCycle = true // ребро назад => цикл
				return
			case white:
				visit(nb)
			}
			// black — ничего: поддерево уже в порядке.
		}
		state[v] = black
		// Вершина «закончена» — пишем в конец, порядок в конце развернём.
		order = append(order, v)
	}

	// Граф может быть несвязным: обходим все вершины.
	for v := 0; v < n && !hasCycle; v++ {
		if state[v] == white {
			visit(v)
		}
	}
	if hasCycle {
		return []int{}
	}
	// Порядок окончания DFS даёт топопорядок в ОБРАТНОМ порядке.
	for i, j := 0, len(order)-1; i < j; i, j = i+1, j-1 {
		order[i], order[j] = order[j], order[i]
	}
	return order
}

// CanFinishDFS — LeetCode 207 через DFS-детекцию цикла.
// prerequisites = [[a, b]] означает ребро b -> a.
func CanFinishDFS(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, len(prerequisites))
	for i, p := range prerequisites {
		edges[i] = []int{p[1], p[0]}
	}
	return len(TopoSortDFS(numCourses, edges)) == numCourses
}

// HasCycleDirected — есть ли цикл в ориентированном графе
// (три состояния, без построения порядка). Отдельно как «мини-шаблон».
func HasCycleDirected(n int, edges [][]int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
	}
	state := make([]int, n) // 0 white, 1 gray, 2 black

	var visit func(v int) bool
	visit = func(v int) bool {
		state[v] = 1
		for _, nb := range graph[v] {
			if state[nb] == 1 {
				return true // цикл
			}
			if state[nb] == 0 && visit(nb) {
				return true
			}
		}
		state[v] = 2
		return false
	}

	for v := 0; v < n; v++ {
		if state[v] == 0 && visit(v) {
			return true
		}
	}
	return false
}
