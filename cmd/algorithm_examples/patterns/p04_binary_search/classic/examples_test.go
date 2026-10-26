package main

import "testing"

func TestBinarySearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11, 13}
	tests := []struct {
		name   string
		target int
		want   int
	}{
		{"в середине", 7, 3},
		{"в начале", 1, 0},
		{"в конце", 13, 6},
		{"не найден меньше", 0, -1},
		{"не найден больше", 100, -1},
		{"не найден между", 4, -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := BinarySearch(nums, tc.target); got != tc.want {
				t.Errorf("BinarySearch(%v, %d) = %d, want %d",
					nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestBinarySearchEdge(t *testing.T) {
	empty := []int{}
	if got := BinarySearch(empty, 1); got != -1 {
		t.Errorf("BinarySearch(empty, 1) = %d, want -1", got)
	}
	single := []int{5}
	if got := BinarySearch(single, 5); got != 0 {
		t.Errorf("BinarySearch([5], 5) = %d, want 0", got)
	}
	if got := BinarySearch(single, 4); got != -1 {
		t.Errorf("BinarySearch([5], 4) = %d, want -1", got)
	}
}

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	tests := []struct {
		target int
		want   int
	}{
		{5, 2},
		{2, 1},
		{7, 4},
		{0, 0},
		{1, 0},
		{6, 3},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := SearchInsert(nums, tc.target); got != tc.want {
				t.Errorf("SearchInsert(%v, %d) = %d, want %d",
					nums, tc.target, got, tc.want)
			}
		})
	}
	if got := SearchInsert([]int{}, 5); got != 0 {
		t.Errorf("SearchInsert(empty, 5) = %d, want 0", got)
	}
}

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		firstBad int
	}{
		{"первая", 5, 1},
		{"последняя", 5, 5},
		{"середина", 5, 3},
		{"большое n", 1000, 500},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			isBad := func(i int) bool { return i >= tc.firstBad }
			if got := FirstBadVersion(tc.n, isBad); got != tc.firstBad {
				t.Errorf("FirstBadVersion(%d, firstBad=%d) = %d, want %d",
					tc.n, tc.firstBad, got, tc.firstBad)
			}
		})
	}
}
