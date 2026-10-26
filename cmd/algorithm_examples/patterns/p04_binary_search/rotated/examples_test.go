package main

import "testing"

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
				t.Errorf("FindMinRotated(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
