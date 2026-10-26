package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"пример из условия", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"все нули", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"нет решений", []int{1, 2, 3}, [][]int{}},
		{"пустой", []int{}, [][]int{}},
		{"меньше трёх", []int{1, 2}, [][]int{}},
		{"много дублей", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ThreeSum(tc.nums)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"пустой", []int{}, nil},
		{"один", []int{1}, []int{1}},
		{"уже отсортирован", []int{0, 1, 2}, []int{0, 1, 2}},
		{"обратный", []int{2, 1, 0}, []int{0, 1, 2}},
		{"смешанный", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"только нули", []int{0, 0, 0}, []int{0, 0, 0}},
		{"только двойки", []int{2, 2, 2}, []int{2, 2, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			SortColors(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}
