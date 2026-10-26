package main

// Паттерн: BFS по графу (соседство, очередь).
//
// Идея: очередь + множество посещённых. Извлекаем узел, смотрим
// всех соседей, непосещённых — помечаем и кладём в очередь.
// Первый раз, когда узел видим, — он и есть кратчайшее расстояние
// от старта (у невзвешенного графа).
//
// Структуры: здесь граф задан списком смежности map[int][]int.
// В задачах LeetCode часто дают срез рёбер — собирайте adjacency
// перед BFS (см. BuildGraph).
//
// Сложность: O(V + E).

// BuildGraph строит список смежности из среза рёбер [from, to].
// Граф неориентированный. Пример рёбер: [][]int{{0,1},{1,2}} ->
// {0:[1], 1:[0,2], 2:[1]}.
func BuildGraph(edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for _, e := range edges {
		from, to := e[0], e[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}

// BFS обходит граф в ширину, возвращая порядок посещения.
// start может отсутствовать в графе — тогда результат пустой.
//
// Пример: 0-1-2, старт 0 -> [0 1 2].
func BFS(graph map[int][]int, start int) []int {
	if _, ok := graph[start]; !ok {
		return []int{}
	}
	visited := map[int]bool{start: true}
	queue := []int{start}
	order := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)

		for _, neighbor := range graph[node] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true // помечаем при постановке
			queue = append(queue, neighbor)
		}
	}
	return order
}

// ShortestPathLength — длина кратчайшего пути от start до end
// в невзвешенном графе (количество рёбер). −1, если пути нет.
//
// Трюк: BFS слой за слоём = расстояние 0, 1, 2, ...
// Как только увидели целевую вершину — вернули её уровень.
//
// Пример: 0-1-2-3, от 0 до 3 -> 3.
func ShortestPathLength(graph map[int][]int, start, end int) int {
	if start == end {
		return 0
	}
	if _, ok := graph[start]; !ok {
		return -1
	}
	visited := map[int]bool{start: true}
	queue := []int{start}
	distance := 0

	for len(queue) > 0 {
		distance++
		for size := len(queue); size > 0; size-- {
			node := queue[0]
			queue = queue[1:]
			for _, neighbor := range graph[node] {
				if visited[neighbor] {
					continue
				}
				if neighbor == end {
					return distance
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return -1
}

// HasPath — существует ли путь между start и end (LeetCode 1971).
// Обход в ширину до первого совпадения; early exit.
//
// Пример: рёбра [[0,1],[1,2],[2,3]], 0->3: true; 0->4: false.
func HasPath(edges [][]int, start, end int) bool {
	graph := BuildGraph(edges)
	if _, ok := graph[start]; !ok {
		return start == end
	}
	visited := map[int]bool{start: true}
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, neighbor := range graph[node] {
			if neighbor == end {
				return true
			}
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			queue = append(queue, neighbor)
		}
	}
	return false
}

// CanVisitAllRooms — можно ли попасть во все комнаты (LeetCode 841).
// Ключи от комнаты i лежат в rooms[i]; из комнаты 0 есть ключи по умолчанию.
//
// Пример: [[1],[2],[3],[]] -> true; [[1,3],[3,0,1],[2],[0]] -> false.
func CanVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	queue := []int{0}
	reached := 1

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]
		for _, key := range rooms[room] {
			if visited[key] {
				continue
			}
			visited[key] = true
			reached++
			queue = append(queue, key)
		}
	}
	return reached == len(rooms)
}
