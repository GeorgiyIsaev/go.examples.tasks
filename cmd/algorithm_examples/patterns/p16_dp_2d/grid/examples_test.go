package main

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m, n int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{1, 1, 1},
		{2, 2, 2},
		{1, 10, 1},
		{10, 1, 1},
		{4, 4, 20},
	}
	for _, tc := range tests {
		if got := UniquePaths(tc.m, tc.n); got != tc.want {
			t.Errorf("UniquePaths(%d, %d) = %d, want %d",
				tc.m, tc.n, got, tc.want)
		}
	}
}

// TestUniquePathsMatchesBrute — рекурсивный перебор всех путей
// на маленьких сетках.
func TestUniquePathsMatchesBrute(t *testing.T) {
	var paths func(i, j, m, n int) int
	paths = func(i, j, m, n int) int {
		if i == m-1 && j == n-1 {
			return 1
		}
		if i >= m || j >= n {
			return 0
		}
		return paths(i+1, j, m, n) + paths(i, j+1, m, n)
	}
	for m := 1; m <= 6; m++ {
		for n := 1; n <= 6; n++ {
			want := paths(0, 0, m, n)
			if got := UniquePaths(m, n); got != want {
				t.Errorf("UniquePaths(%d, %d) = %d, brute = %d", m, n, got, want)
			}
		}
	}
}

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"классический пример", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"второй пример LC", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{"одна клетка", [][]int{{5}}, 5},
		{"только вправо", [][]int{{1, 2, 3}}, 6},
		{"только вниз", [][]int{{1}, {2}, {3}}, 6},
		{"все нули", [][]int{{0, 0}, {0, 0}}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinPathSum(tc.grid); got != tc.want {
				t.Errorf("MinPathSum(%v) = %d, want %d", tc.grid, got, tc.want)
			}
		})
	}
}

// TestMinPathSumMatchesBrute — перебор двух путей на каждой клетке.
func TestMinPathSumMatchesBrute(t *testing.T) {
	var best func(grid [][]int, i, j int) int
	best = func(grid [][]int, i, j int) int {
		m, n := len(grid), len(grid[0])
		if i == m-1 && j == n-1 {
			return grid[i][j]
		}
		down := int(1 << 30)
		right := int(1 << 30)
		if i+1 < m {
			down = best(grid, i+1, j)
		}
		if j+1 < n {
			right = best(grid, i, j+1)
		}
		if down < right {
			return grid[i][j] + down
		}
		return grid[i][j] + right
	}

	cases := [][][]int{
		{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		{{1, 2, 3}, {4, 5, 6}},
		{{9, 1, 1}, {1, 9, 1}, {1, 1, 9}},
	}
	for _, grid := range cases {
		want := best(grid, 0, 0)
		if got := MinPathSum(grid); got != want {
			t.Errorf("MinPathSum(%v) = %d, brute = %d", grid, got, want)
		}
	}
}
