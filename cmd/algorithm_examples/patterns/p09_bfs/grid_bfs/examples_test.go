package main

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"два острова", [][]int{
			{1, 1, 0},
			{1, 0, 0},
			{0, 0, 1},
		}, 2},
		{"все суша — один остров", [][]int{
			{1, 1},
			{1, 1},
		}, 1},
		{"диагонали не соединены", [][]int{
			{1, 0},
			{0, 1},
		}, 2},
		{"все вода", [][]int{
			{0, 0},
			{0, 0},
		}, 0},
		{"одна клетка суша", [][]int{{1}}, 1},
		{"одна клетка вода", [][]int{{0}}, 0},
		{"пустая сетка", [][]int{}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NumIslands(tc.grid); got != tc.want {
				t.Errorf("NumIslands() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestRottingOranges(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"классический пример", [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}, 4},
		{"недостижимый апельсин", [][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}, -1},
		{"уже все гнилые", [][]int{
			{2, 2},
			{2, 2},
		}, 0},
		{"нет свежих", [][]int{
			{0, 2},
			{0, 0},
		}, 0},
		{"один свежий рядом", [][]int{{2, 1}}, 1},
		{"только свежие", [][]int{{1, 1}, {1, 1}}, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RottingOranges(tc.grid); got != tc.want {
				t.Errorf("RottingOranges() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"классический пример", [][]int{
			{0, 0, 0},
			{1, 1, 0},
			{1, 1, 0},
		}, 4},
		{"конечная клетка — стена", [][]int{
			{0, 1},
			{1, 1},
		}, -1},
		{"диагональ разрешена — путь есть", [][]int{
			{0, 1},
			{1, 0},
		}, 2},
		{"старт закрыт", [][]int{
			{1, 0},
			{0, 0},
		}, -1},
		{"1x1 свободна", [][]int{{0}}, 1},
		{"1x1 стена", [][]int{{1}}, -1},
		{"по диагонали", [][]int{
			{0, 0},
			{0, 0},
		}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShortestPathBinaryMatrix(tc.grid); got != tc.want {
				t.Errorf("ShortestPathBinaryMatrix() = %d, want %d", got, tc.want)
			}
		})
	}
}
