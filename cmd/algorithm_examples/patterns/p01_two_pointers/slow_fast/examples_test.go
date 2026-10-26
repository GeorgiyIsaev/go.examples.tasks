package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicatesSorted(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"пустой", []int{}, 0},
		{"один", []int{1}, 1},
		{"все одинаковые", []int{2, 2, 2, 2}, 1},
		{"все уникальные", []int{1, 2, 3, 4}, 4},
		{"с дублями", []int{1, 1, 2, 3, 3, 3, 4}, 4},
		{"дубли в конце", []int{1, 2, 3, 3, 3}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			got := RemoveDuplicatesSorted(nums)
			if got != tc.want {
				t.Errorf("RemoveDuplicatesSorted(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"пустой", []int{}, nil},
		{"без нулей", []int{1, 2, 3}, []int{1, 2, 3}},
		{"только нули", []int{0, 0, 0}, []int{0, 0, 0}},
		{"чередование", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"нули в конце", []int{1, 2, 3, 0, 0}, []int{1, 2, 3, 0, 0}},
		{"нули в начале", []int{0, 0, 1, 2, 3}, []int{1, 2, 3, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			MoveZeroes(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Errorf("MoveZeroes(%v) = %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{"пустой", []int{}, 1, 0},
		{"нет совпадений", []int{1, 2, 3}, 4, 3},
		{"всё совпадает", []int{3, 3, 3}, 3, 0},
		{"частично", []int{3, 2, 2, 3}, 3, 2},
		{"в середине", []int{1, 2, 3, 4, 2}, 2, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			got := RemoveElement(nums, tc.val)
			if got != tc.want {
				t.Errorf("RemoveElement(%v, %d) = %d, want %d", tc.nums, tc.val, got, tc.want)
			}
		})
	}
}
