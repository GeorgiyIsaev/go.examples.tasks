package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"классический пример", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"второй пример LC", []int{4, 2, 0, 3, 2, 5}, 9},
		{"без ям", []int{1, 2, 3, 4}, 0},
		{"без ям (убывание)", []int{4, 3, 2, 1}, 0},
		{"один столбик", []int{5}, 0},
		{"ровная плоскость", []int{2, 2, 2}, 0},
		{"впадина по краям", []int{3, 0, 3}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Trap(tc.heights); got != tc.want {
				t.Errorf("Trap(%v) = %d, want %d", tc.heights, got, tc.want)
			}
		})
	}
}

func TestTrapMatchesTwoPointers(t *testing.T) {
	// Эталон: классический двухуказательный алгоритм O(1) памяти.
	brute := func(h []int) int {
		left, right := 0, len(h)-1
		leftMax, rightMax := 0, 0
		water := 0
		for left < right {
			if h[left] < h[right] {
				if h[left] >= leftMax {
					leftMax = h[left]
				} else {
					water += leftMax - h[left]
				}
				left++
			} else {
				if h[right] >= rightMax {
					rightMax = h[right]
				} else {
					water += rightMax - h[right]
				}
				right--
			}
		}
		return water
	}

	cases := [][]int{
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		{4, 2, 0, 3, 2, 5},
		{5, 4, 3, 2, 1, 2},
		{2, 2, 2},
	}
	for _, h := range cases {
		want := brute(h)
		if got := Trap(h); got != want {
			t.Errorf("Trap(%v) = %d, two pointers = %d", h, got, want)
		}
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"классический пример", []int{2, 1, 5, 6, 2, 3}, 10},
		{"два столбика", []int{2, 4}, 4},
		{"один столбик", []int{1}, 1},
		{"возрастающие", []int{1, 2, 3}, 4}, // 2×2
		{"убывающие", []int{3, 2, 1}, 4},    // 2×2
		{"нули", []int{0, 0, 0}, 0},
		{"все равные", []int{3, 3, 3, 3}, 12},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LargestRectangleArea(tc.heights); got != tc.want {
				t.Errorf("LargestRectangleArea(%v) = %d, want %d",
					tc.heights, got, tc.want)
			}
		})
	}
}

func TestLargestRectangleAreaBruteForce(t *testing.T) {
	// Эталон: для каждой пары границ — минимум высот × ширина.
	brute := func(h []int) int {
		best := 0
		for i := 0; i < len(h); i++ {
			m := h[i]
			for j := i; j < len(h); j++ {
				if h[j] < m {
					m = h[j]
				}
				if area := m * (j - i + 1); area > best {
					best = area
				}
			}
		}
		return best
	}

	cases := [][]int{
		{2, 1, 5, 6, 2, 3},
		{2, 4},
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{0, 1, 0, 1},
	}
	for _, h := range cases {
		want := brute(h)
		if got := LargestRectangleArea(h); got != want {
			t.Errorf("LargestRectangleArea(%v) = %d, brute force = %d", h, got, want)
		}
	}
}
