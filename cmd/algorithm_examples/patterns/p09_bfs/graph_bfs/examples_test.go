package main

import (
	"reflect"
	"testing"
)

func TestBuildGraph(t *testing.T) {
	g := BuildGraph([][]int{{0, 1}, {1, 2}})
	want := map[int][]int{
		0: {1},
		1: {0, 2},
		2: {1},
	}
	if !reflect.DeepEqual(g, want) {
		t.Errorf("BuildGraph() = %v, want %v", g, want)
	}
	if len(BuildGraph(nil)) != 0 {
		t.Errorf("BuildGraph(nil) should be empty")
	}
}

func TestBFS(t *testing.T) {
	graph := BuildGraph([][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}})
	got := BFS(graph, 0)

	if len(got) != 4 {
		t.Fatalf("BFS() visited %d nodes, want 4: %v", len(got), got)
	}
	if got[0] != 0 {
		t.Errorf("BFS() must start at 0, got %v", got)
	}
	seen := map[int]bool{}
	for _, n := range got {
		if seen[n] {
			t.Errorf("node %d visited twice: %v", n, got)
		}
		seen[n] = true
	}

	if order := BFS(graph, 99); len(order) != 0 {
		t.Errorf("BFS from missing node = %v, want []", order)
	}
}

func TestShortestPathLength(t *testing.T) {
	// линия 0-1-2-3
	line := BuildGraph([][]int{{0, 1}, {1, 2}, {2, 3}})
	// квадрат 0-1, 1-2, 2-3, 3-0 + хвост 1-4
	g := BuildGraph([][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {1, 4}})

	tests := []struct {
		name       string
		graph      map[int][]int
		start, end int
		want       int
	}{
		{"туда и обратно", line, 0, 3, 3},
		{"start == end", line, 2, 2, 0},
		{"через квадрат короче", g, 0, 2, 2},
		{"до хвоста", g, 0, 4, 2},
		{"нет пути", line, 0, 99, -1},
		{"неизолированная вершина", line, 99, 0, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShortestPathLength(tc.graph, tc.start, tc.end); got != tc.want {
				t.Errorf("ShortestPathLength(%d, %d) = %d, want %d",
					tc.start, tc.end, got, tc.want)
			}
		})
	}
}

func TestHasPath(t *testing.T) {
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}}
	tests := []struct {
		name       string
		start, end int
		want       bool
	}{
		{"существует", 0, 3, true},
		{"один и тот же узел", 5, 5, true},
		{"нет узла", 0, 9, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasPath(edges, tc.start, tc.end); got != tc.want {
				t.Errorf("HasPath(%d, %d) = %v, want %v",
					tc.start, tc.end, got, tc.want)
			}
		})
	}

	if HasPath([][]int{{0, 1}, {2, 3}}, 0, 3) {
		t.Errorf("HasPath across components should be false")
	}
}

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  bool
	}{
		{"все достижимы", [][]int{{1}, {2}, {3}, {}}, true},
		{"не все", [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
		{"одна комната", [][]int{{}}, true},
		{"две комнаты без ключа", [][]int{{}, {0}}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanVisitAllRooms(tc.rooms); got != tc.want {
				t.Errorf("CanVisitAllRooms(%v) = %v, want %v",
					tc.rooms, got, tc.want)
			}
		})
	}
}
