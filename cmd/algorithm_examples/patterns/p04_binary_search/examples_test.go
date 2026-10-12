package main

import (
	"reflect"
	"testing"
)

// ===== Classic =====

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

// ===== Bounds =====

func TestLowerBound(t *testing.T) {
	nums := []int{1, 3, 3, 5, 7, 7, 7, 9}
	tests := []struct {
		target int
		want   int
	}{
		{0, 0},
		{1, 0},
		{3, 1},
		{5, 3},
		{7, 4},
		{9, 7},
		{10, 8},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := LowerBound(nums, tc.target); got != tc.want {
				t.Errorf("LowerBound(%v, %d) = %d, want %d",
					nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	nums := []int{1, 3, 3, 5, 7, 7, 7, 9}
	tests := []struct {
		target int
		want   int
	}{
		{0, 0},
		{1, 1},
		{3, 3},
		{5, 4},
		{7, 7},
		{9, 8},
		{10, 8},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := UpperBound(nums, tc.target); got != tc.want {
				t.Errorf("UpperBound(%v, %d) = %d, want %d",
					nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"пример", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"не найден", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"все одинаковые", []int{5, 5, 5}, 5, []int{0, 2}},
		{"один элемент найден", []int{5}, 5, []int{0, 0}},
		{"один элемент не найден", []int{5}, 4, []int{-1, -1}},
		{"пустой", []int{}, 5, []int{-1, -1}},
		{"в начале", []int{5, 5, 7, 8}, 5, []int{0, 1}},
		{"в конце", []int{1, 2, 5, 5}, 5, []int{2, 3}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SearchRange(tc.nums, tc.target); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SearchRange(%v, %d) = %v, want %v",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestCountOccurrences(t *testing.T) {
	nums := []int{1, 3, 3, 5, 7, 7, 7, 9}
	tests := []struct {
		target int
		want   int
	}{
		{1, 1},
		{3, 2},
		{7, 3},
		{9, 1},
		{4, 0},
		{0, 0},
		{10, 0},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountOccurrences(nums, tc.target); got != tc.want {
				t.Errorf("CountOccurrences(%v, %d) = %d, want %d",
					nums, tc.target, got, tc.want)
			}
		})
	}
}

// ===== Rotated =====

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"пример, найдено в правой", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"пример, найдено в левой", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{"не найдено", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"не повернут", []int{1, 2, 3, 4, 5}, 3, 2},
		{"один элемент найден", []int{1}, 1, 0},
		{"один элемент не найден", []int{1}, 0, -1},
		{"два элемента, найдено", []int{3, 1}, 1, 1},
		{"первый элемент", []int{4, 5, 6, 7, 0, 1, 2}, 4, 0},
		{"последний элемент", []int{4, 5, 6, 7, 0, 1, 2}, 2, 6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SearchRotated(tc.nums, tc.target); got != tc.want {
				t.Errorf("SearchRotated(%v, %d) = %d, want %d",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestFindMinRotated(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"пример", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"не повернут", []int{1, 2, 3, 4, 5}, 1},
		{"один", []int{1}, 1},
		{"два, повернут", []int{2, 1}, 1},
		{"два, не повернут", []int{1, 2}, 1},
		{"последний минимум", []int{5, 1, 2, 3, 4}, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindMinRotated(tc.nums); got != tc.want {
				t.Errorf("FindMinRotated(%v) = %d, want %d",
					tc.nums, got, tc.want)
			}
		})
	}
}

// ===== Binary Search on Answer =====

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 2},
		{9, 3},
		{15, 3},
		{16, 4},
		{100, 10},
		{1_000_000, 1000},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := MySqrt(tc.x); got != tc.want {
				t.Errorf("MySqrt(%d) = %d, want %d", tc.x, got, tc.want)
			}
		})
	}
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"пример 1", []int{3, 6, 7, 11}, 8, 4},
		{"пример 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"пример 3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"одна куча", []int{10}, 1, 10},
		{"одна куча, много часов", []int{10}, 10, 1},
		{"все единицы", []int{1, 1, 1, 1}, 4, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinEatingSpeed(tc.piles, tc.h); got != tc.want {
				t.Errorf("MinEatingSpeed(%v, %d) = %d, want %d",
					tc.piles, tc.h, got, tc.want)
			}
		})
	}
}

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		days    int
		want    int
	}{
		{"пример", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{"пример 2", []int{3, 2, 2, 4, 1, 4}, 3, 6},
		{"пример 3", []int{1, 2, 3, 1, 1}, 4, 3},
		{"один груз", []int{5}, 1, 5},
		{"много дней", []int{1, 2, 3}, 10, 3},
		{"один день", []int{1, 2, 3}, 1, 6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShipWithinDays(tc.weights, tc.days); got != tc.want {
				t.Errorf("ShipWithinDays(%v, %d) = %d, want %d",
					tc.weights, tc.days, got, tc.want)
			}
		})
	}
}
