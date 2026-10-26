package main

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		{"классический пример", []int{3, 9, 20, -1, -1, 15, 7},
			[][]int{{3}, {9, 20}, {15, 7}}},
		{"один узел", []int{1}, [][]int{{1}}},
		{"пустое дерево", nil, [][]int{}},
		{"все левые", []int{1, 2, -1, 3}, [][]int{{1}, {2}, {3}}},
		{"полное до 2 уровня", []int{1, 2, 3, 4, 5, 6, 7},
			[][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LevelOrder(BuildTree(tc.vals))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LevelOrder(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func TestAverageOfLevels(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []float64
	}{
		{"классический пример", []int{3, 9, 20, -1, -1, 15, 7},
			[]float64{3, 14.5, 11}},
		{"один узел", []int{1}, []float64{1}},
		{"пустое дерево", nil, []float64{}},
		{"два уровня", []int{1, 2, 3}, []float64{1, 2.5}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AverageOfLevels(BuildTree(tc.vals))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("AverageOfLevels(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"классический пример", []int{3, 9, 20, -1, -1, 15, 7}, 2},
		{"один узел", []int{1}, 1},
		{"пустое дерево", nil, 0},
		{"кривое дерево", []int{1, 2, -1, 3, 4}, 3},
		{"лево короче", []int{1, 2, 3, 4}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinDepth(BuildTree(tc.vals)); got != tc.want {
				t.Errorf("MinDepth(%v) = %d, want %d", tc.vals, got, tc.want)
			}
		})
	}
}

func TestMaxLevelSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"второй уровень максимален", []int{1, 7, 0, 7, -8, -1, -1}, 2},
		{"корень максимален", []int{9, 1, 0, 0, 0}, 1},
		{"один узел", []int{42}, 1},
		{"равные суммы — берём меньший уровень", []int{5, 2, 3}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxLevelSum(BuildTree(tc.vals)); got != tc.want {
				t.Errorf("MaxLevelSum(%v) = %d, want %d", tc.vals, got, tc.want)
			}
		})
	}
}
