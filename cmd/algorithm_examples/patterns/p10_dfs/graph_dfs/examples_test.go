package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestDFS(t *testing.T) {
	graph := BuildGraph([][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}})

	got := DFS(graph, 0)
	if len(got) != 4 {
		t.Fatalf("DFS() visited %d nodes, want 4: %v", len(got), got)
	}
	if got[0] != 0 {
		t.Errorf("DFS() must start at 0, got %v", got)
	}
	sorted := append([]int{}, got...)
	sort.Ints(sorted)
	if want := []int{0, 1, 2, 3}; !reflect.DeepEqual(sorted, want) {
		t.Errorf("DFS() visited %v, want %v", sorted, want)
	}

	if order := DFS(graph, 99); len(order) != 0 {
		t.Errorf("DFS from missing node = %v, want []", order)
	}
}

func TestDFSIterativeMatchesRecursive(t *testing.T) {
	graphs := []map[int][]int{
		BuildGraph([][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}),
		BuildGraph([][]int{{0, 1}, {1, 2}, {2, 3}}),
		BuildGraph([][]int{{0, 1}}),
	}

	for i, g := range graphs {
		rec := DFS(g, 0)
		iter := DFSIterative(g, 0)
		sort.Ints(rec)
		sort.Ints(iter)
		if !reflect.DeepEqual(rec, iter) {
			t.Errorf("graph %d: iterative %v != recursive %v", i, iter, rec)
		}
	}
}

func TestHasPathDFS(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	tests := []struct {
		name       string
		start, end int
		want       bool
	}{
		{"существует", 0, 3, true},
		{"тот же узел", 5, 5, true},
		{"нет узла", 0, 9, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasPathDFS(edges, tc.start, tc.end); got != tc.want {
				t.Errorf("HasPathDFS(%d, %d) = %v, want %v",
					tc.start, tc.end, got, tc.want)
			}
		})
	}

	if HasPathDFS([][]int{{0, 1}, {2, 3}}, 0, 3) {
		t.Errorf("HasPathDFS across components should be false")
	}
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  bool
	}{
		{"квадрат — цикл", [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, true},
		{"линия — без цикла", [][]int{{0, 1}, {1, 2}}, false},
		{"дерево — без цикла", [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}, false},
		{"цикл во второй компоненте", [][]int{{0, 1}, {2, 3}, {3, 4}, {4, 2}}, true},
		{"нет рёбер", nil, false},
		{"самопетля", [][]int{{0, 0}}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasCycle(tc.edges); got != tc.want {
				t.Errorf("HasCycle(%v) = %v, want %v", tc.edges, got, tc.want)
			}
		})
	}
}

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"две компоненты", 4, [][]int{{0, 1}, {2, 3}}, 2},
		{"одна компонента и один одиночка", 4, [][]int{{0, 1}, {1, 2}}, 2},
		{"все связаны", 3, [][]int{{0, 1}, {1, 2}, {0, 2}}, 1},
		{"без рёбер", 3, nil, 3},
		{"ноль вершин", 0, nil, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountComponents(tc.n, tc.edges); got != tc.want {
				t.Errorf("CountComponents(%d, %v) = %d, want %d",
					tc.n, tc.edges, got, tc.want)
			}
		})
	}
}
