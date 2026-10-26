package main

// Паттерн: топологическая сортировка — алгоритм Кана (Kahn).
//
// Идея: в ориентированном графе без циклов (DAG) всегда есть вершина
// без входящих рёбер (indegree 0). Забираем её в начало порядка,
// «снимаем» её рёбра (уменьшаем indegree соседей) — и повторяем.
// Полученный порядок гарантирует: **все рёбра смотрят вперёд** —
// если a -> b, то a стоит раньше b.
//
// Если граф содержит цикл — в какой-то момент не останется вершин
// с indegree 0, а в порядке окажется меньше n вершин. Это и есть
// детектор цикла (главное в задачах LC207/LC210).
//
// Сложность: O(V + E) — каждая вершина и каждое ребро обрабатываются
// ровно один раз.
//
// Словарь: graph[a] = [b, c] означает рёбра a->b, a->c
// (a — предшественник, b — преемник).

// TopoSortKahn — топологический порядок вершин 0..n-1.
// Если цикл — возвращает пустой срез (порядка не существует).
//
// Пример: рёбра [[0,1],[0,2],[1,3],[2,3]] -> [0 1 2 3] (или [0 2 1 3]).
func TopoSortKahn(n int, edges [][]int) []int {
	// Список смежности и входящая степень каждой вершины.
	graph := make([][]int, n)
	indegree := make([]int, n)
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
		indegree[to]++
	}

	// В очередь — все вершины без входящих рёбер.
	queue := []int{}
	for v := 0; v < n; v++ {
		if indegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	order := make([]int, 0, n)
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		order = append(order, v)

		for _, nb := range graph[v] {
			indegree[nb]--
			if indegree[nb] == 0 {
				queue = append(queue, nb)
			}
		}
	}

	// Обошли не все вершины -> есть цикл.
	if len(order) != n {
		return []int{}
	}
	return order
}

// CanFinish — можно ли закончить все курсы (LeetCode 207):
// есть ли цикл в графе предшественств (prerequisites [a,b] = b раньше a).
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Переворачиваем: [b,a] означает ребро b->a? Нет: prerequisites[i] =
	// [a, b] значит «a зависит от b», т.е. ребро b -> a.
	edges := make([][]int, len(prerequisites))
	for i, p := range prerequisites {
		edges[i] = []int{p[1], p[0]}
	}
	return len(TopoSortKahn(numCourses, edges)) == numCourses
}

// FindOrder — кратчайший (?) валидный порядок прохождения курсов
// (LeetCode 210). Любой корректный порядок подходит; при цикле — [].
//
// Пример: n=4, prerequisites [[1,0],[2,0],[3,1],[3,2]] -> [0 1 2 3]
// (или [0 2 1 3]).
func FindOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, len(prerequisites))
	for i, p := range prerequisites {
		edges[i] = []int{p[1], p[0]}
	}
	return TopoSortKahn(numCourses, edges)
}
