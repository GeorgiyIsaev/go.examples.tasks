package main

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5}, []int{5}},
		{[]int{}, []int{}},
		{[]int{-1, 2, -3}, []int{-1, 1, -2}},
	}
	for _, tc := range tests {
		got := RunningSum(tc.nums)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("RunningSum(%v) = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{1, 7, 3, 6, 5, 6}, 3},
		{"нет решения", []int{1, 2, 3}, -1},
		{"первый индекс", []int{2, -1, 1}, 0},
		{"ни один не подходит", []int{-1, -1, 2}, -1},
		{"ноль посередине", []int{0, 0, 0}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := PivotIndex(tc.nums); got != tc.want {
				t.Errorf("PivotIndex(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestPivotIndexMatchesBrute(t *testing.T) {
	cases := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
		{0, 0, 0, 1},
		{-5, 2, 3, -2},
	}
	for _, nums := range cases {
		want := -1
		for i := range nums {
			left, right := 0, 0
			for j := 0; j < i; j++ {
				left += nums[j]
			}
			for j := i + 1; j < len(nums); j++ {
				right += nums[j]
			}
			if left == right {
				want = i
				break
			}
		}
		if got := PivotIndex(nums); got != want {
			t.Errorf("PivotIndex(%v) = %d, brute = %d", nums, got, want)
		}
	}
}

func TestRangeSumQuery(t *testing.T) {
	q := NewRangeSumQuery([]int{-2, 0, 3, -5, 2, -1})
	tests := []struct {
		left, right int
		want        int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
		{3, 3, -5},
	}
	for _, tc := range tests {
		if got := q.SumRange(tc.left, tc.right); got != tc.want {
			t.Errorf("SumRange(%d,%d) = %d, want %d",
				tc.left, tc.right, got, tc.want)
		}
	}

	// Сверка с лобовым суммированием по всем парам.
	nums := []int{-2, 0, 3, -5, 2, -1}
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if got := q.SumRange(i, j); got != sum {
				t.Errorf("SumRange(%d,%d) = %d, brute = %d", i, j, got, sum)
			}
		}
	}
}
