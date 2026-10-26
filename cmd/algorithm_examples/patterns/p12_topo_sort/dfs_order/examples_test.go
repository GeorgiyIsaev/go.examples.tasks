package main

import "testing"

// validOrder — свойство топопорядка: для каждого ребра a->b
// a встречается раньше b, каждая вершина ровно один раз.
func validOrder(order []int, n int, edges [][]int) bool {
	if len(order) != n {
		return false
	}
	pos := make(map[int]int, n)
	for i, v := range order {
		if _, seen := pos[v]; seen {
			return false
		}
		pos[v] = i
	}
	for _, e := range edges {
		if pos[e[0]] >= pos[e[1]] {
			return false
		}
	}
	return true
}

func TestTopoSortDFS(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
	}{
		{"классический пример", 4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
		{"линия", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{"без рёбер", 3, nil},
		{"развязанные компоненты", 4, [][]int{{0, 1}, {2, 3}}},
		{"полный DAG", 3, [][]int{{0, 1}, {0, 2}, {1, 2}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopoSortDFS(tc.n, tc.edges)
			if !validOrder(got, tc.n, tc.edges) {
				t.Errorf("TopoSortDFS(%d, %v) = %v, not a valid topo order",
					tc.n, tc.edges, got)
			}
		})
	}
}

func TestTopoSortDFSCycle(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
	}{
		{"простой цикл", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}},
		{"самопетля", 2, [][]int{{0, 0}}},
		{"цикл во второй компоненте", 4, [][]int{{0, 1}, {2, 3}, {3, 2}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := TopoSortDFS(tc.n, tc.edges); len(got) != 0 {
				t.Errorf("TopoSortDFS with cycle = %v, want []", got)
			}
		})
	}
}

func TestCanFinishDFS(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		prereq [][]int
		want   bool
	}{
		{"можно", 2, [][]int{{1, 0}}, true},
		{"цикл", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"без предпосылок", 3, nil, true},
		{"длинный путь", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{"трицикл", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanFinishDFS(tc.n, tc.prereq); got != tc.want {
				t.Errorf("CanFinishDFS(%d, %v) = %v, want %v",
					tc.n, tc.prereq, got, tc.want)
			}
		})
	}
}

func TestHasCycleDirected(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{"цепочка — без цикла", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}, false},
		{"цикл", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, true},
		{"нет рёбер", 3, nil, false},
		{"две компоненты, цикл во второй", 5, [][]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true},
		{"ребро назад в DAG — не цикл", 3, [][]int{{0, 2}, {1, 2}}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasCycleDirected(tc.n, tc.edges); got != tc.want {
				t.Errorf("HasCycleDirected(%d, %v) = %v, want %v",
					tc.n, tc.edges, got, tc.want)
			}
		})
	}
}

// kahnVerdict — упрощённый алгоритм Кана прямо в тесте: служит
// независимой второй реализацией для перекрёстной проверки DFS.
func kahnVerdict(n int, edges [][]int) bool {
	graph := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		indeg[e[1]]++
	}
	queue := []int{}
	for v := 0; v < n; v++ {
		if indeg[v] == 0 {
			queue = append(queue, v)
		}
	}
	taken := 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		taken++
		for _, nb := range graph[v] {
			indeg[nb]--
			if indeg[nb] == 0 {
				queue = append(queue, nb)
			}
		}
	}
	return taken == n
}

// TestKahnAndDFSSameVerdict — обе реализации обязаны согласовываться:
// если одна нашла порядок, другая тоже находит.
func TestKahnAndDFSSameVerdict(t *testing.T) {
	cases := []struct {
		n     int
		edges [][]int
	}{
		{4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
		{4, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{3, [][]int{{0, 1}, {1, 2}, {2, 0}}},
		{3, nil},
	}
	for _, tc := range cases {
		kahn := kahnVerdict(tc.n, tc.edges)
		dfs := len(TopoSortDFS(tc.n, tc.edges)) > 0
		if kahn != dfs {
			t.Errorf("n=%d edges=%v: kahn found order=%v, dfs found order=%v",
				tc.n, tc.edges, kahn, dfs)
		}
	}
}
