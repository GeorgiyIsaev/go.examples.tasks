package main

import (
	"reflect"
	"testing"
)

// validOrder проверяет свойство топопорядка: для каждого ребра a->b
// индекс a в порядке меньше индекса b. И что порядок содержит все n
// вершин ровно по разу.
func validOrder(order []int, n int, edges [][]int) bool {
	if len(order) != n {
		return false
	}
	pos := make(map[int]int, n)
	for i, v := range order {
		if _, seen := pos[v]; seen {
			return false // вершина повторилась
		}
		pos[v] = i
	}
	for _, e := range edges {
		if pos[e[0]] >= pos[e[1]] {
			return false // ребро «назад»
		}
	}
	return true
}

func TestTopoSortKahn(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
	}{
		{"классический пример", 4, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}},
		{"линия", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{"без рёбер", 3, nil},
		{"полный DAG", 3, [][]int{{0, 1}, {0, 2}, {1, 2}}},
		{"развязанные компоненты", 4, [][]int{{0, 1}, {2, 3}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopoSortKahn(tc.n, tc.edges)
			if !validOrder(got, tc.n, tc.edges) {
				t.Errorf("TopoSortKahn(%d, %v) = %v, not a valid topo order",
					tc.n, tc.edges, got)
			}
		})
	}
}

func TestTopoSortKahnCycle(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
	}{
		{"простой цикл", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}},
		{"цикл в одной компоненте", 4, [][]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}}},
		{"самопетля", 2, [][]int{{0, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := TopoSortKahn(tc.n, tc.edges); len(got) != 0 {
				t.Errorf("TopoSortKahn with cycle = %v, want []", got)
			}
		})
	}
}

func TestCanFinish(t *testing.T) {
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
			if got := CanFinish(tc.n, tc.prereq); got != tc.want {
				t.Errorf("CanFinish(%d, %v) = %v, want %v",
					tc.n, tc.prereq, got, tc.want)
			}
		})
	}
}

func TestFindOrder(t *testing.T) {
	// Общие случаи: порядок любой валидный — проверяем свойство.
	prereq := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}} // развёрнутые
	if got := FindOrder(4, prereq); !validOrder(got, 4, edges) {
		t.Errorf("FindOrder(4, ...) = %v, not a valid order", got)
	}

	// Нет курсов — пусто.
	if got := FindOrder(0, nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("FindOrder(0, nil) = %v, want []", got)
	}

	// Цикл -> [].
	if got := FindOrder(2, [][]int{{1, 0}, {0, 1}}); len(got) != 0 {
		t.Errorf("FindOrder with cycle = %v, want []", got)
	}
}
