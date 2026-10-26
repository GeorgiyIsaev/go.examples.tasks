package main

import (
	"reflect"
	"testing"
)

func TestNumIslandsDFS(t *testing.T) {
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
		{"диагонали не соединены", [][]int{
			{1, 0},
			{0, 1},
		}, 2},
		{"все вода", [][]int{{0, 0}, {0, 0}}, 0},
		{"одна клетка", [][]int{{1}}, 1},
		{"пустая сетка", [][]int{}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NumIslandsDFS(tc.grid); got != tc.want {
				t.Errorf("NumIslandsDFS() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"все 1 связны в одну восьмёрку", [][]int{
			{0, 0, 1},
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
		}, 8},
		{"все нули", [][]int{
			{0, 0},
			{0, 0},
		}, 0},
		{"один остров 4 клетки", [][]int{
			{1, 1},
			{1, 1},
		}, 4},
		{"два острова 3 и 1", [][]int{
			{1, 1, 1, 0, 1},
		}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxAreaOfIsland(tc.grid); got != tc.want {
				t.Errorf("MaxAreaOfIsland() = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name  string
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{"классический пример", [][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		}, 1, 1, 2, [][]int{
			{2, 2, 2},
			{2, 2, 0},
			{2, 0, 1},
		}},
		{"цвет уже совпадает", [][]int{
			{0, 0, 0},
			{0, 0, 0},
		}, 0, 0, 0, [][]int{
			{0, 0, 0},
			{0, 0, 0},
		}},
		{"диагональ не заливается", [][]int{
			{1, 0},
			{0, 1},
		}, 0, 0, 9, [][]int{
			{9, 0},
			{0, 1},
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FloodFill(tc.image, tc.sr, tc.sc, tc.color)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FloodFill() = %v, want %v", got, tc.want)
			}
		})
	}
}
