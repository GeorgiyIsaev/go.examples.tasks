package main

// Паттерн: топологическая сортировка как основа для запросов
// «является ли u предшественником v» (LeetCode 1462 Course Schedule IV).
//
// Задача: рёбра зависимостей (prerequisites[i] = [a, b] — «a идёт до b»)
// и запросы [u, v]: существует ли путь u -> ... -> v?
//
// Два подхода, оба здесь:
//
// 1. **Запрос за запросом** — DFS достижимости из u под каждый вопрос:
//    O(Q · (V + E)). Просто, работает, когда запросов мало.
// 2. **Предвычисление** — DFS от каждой вершины один раз, готовая
//    таблица reachable[u][v]: O(V · (V + E)) времени и O(V²) памяти,
//    каждый запрос — O(1). Выигрывает, когда Q ~ V².
//
// Общее: граф предшественств обычно DAG. Цикл здесь не запрещён —
// он просто сделал бы участников цикла взаимно достижимыми, — но
// смысл задачи («порядок курсов») предполагает DAG.
// Сложность обхода O(V + E), как у всех обходов.

// CheckIfPrerequisite — под каждый запрос отдельный DFS.
// prerequisites = [[a, b]] означает ребро a -> b,
// queries[i] = [u, v] — вопрос «u предшественник v?».
//
// Пример: n=4, prereq=[[1,0],[2,0],[3,1],[3,2]],
// query [3,0] -> true (3->1->0), [0,3] -> false.
func CheckIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	graph := make([][]int, n)
	for _, e := range prerequisites {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	result := make([]bool, len(queries))
	for qi, q := range queries {
		src, dst := q[0], q[1]
		// DFS достижимости из src.
		visited := make([]bool, n)
		stack := []int{src}
		visited[src] = true
		found := false
		for len(stack) > 0 && !found {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, nb := range graph[v] {
				if nb == dst {
					found = true
					break
				}
				if visited[nb] {
					continue
				}
				visited[nb] = true
				stack = append(stack, nb)
			}
		}
		result[qi] = found
	}
	return result
}

// PredecessorQueries — предвычисляем таблицу достижимости от каждой
// вершины, затем отвечаем на все запросы за O(1).
//
// Таблица reachable[u][v] = «с u можно добраться до v».
// Семантика строгая: u не предшественник самого себя — по умолчанию
// reachable[u][u] = false; станет true только при цикле через u
// (тогда путь u -> ... -> u существует). Так же считает и DFS-вариант.
func PredecessorQueries(n int, prerequisites [][]int, queries [][]int) []bool {
	graph := make([][]int, n)
	for _, e := range prerequisites {
		graph[e[0]] = append(graph[e[0]], e[1])
	}

	// DFS от каждой вершины: O(V · (V + E)).
	reachable := make([][]bool, n)
	for u := 0; u < n; u++ {
		reachable[u] = make([]bool, n)
		stack := []int{u}
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, nb := range graph[v] {
				if reachable[u][nb] {
					continue
				}
				reachable[u][nb] = true
				stack = append(stack, nb)
			}
		}
	}

	result := make([]bool, len(queries))
	for i, q := range queries {
		result[i] = reachable[q[0]][q[1]]
	}
	return result
}
