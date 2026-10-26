package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"классический пример", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"ответ не в начале", []int{3, 2, 4}, 6, []int{1, 2}},
		{"дубли значений", []int{3, 3}, 6, []int{0, 1}},
		{"с отрицательными", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}

	if got := TwoSum([]int{1, 2, 3}, 100); len(got) != 0 {
		t.Errorf("TwoSum without pair = %v, want []", got)
	}
}

func TestTwoSumCount(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"все дубли", []int{1, 1, 1}, 2, 3},
		{"одна пара", []int{1, 2, 3}, 3, 1},
		{"без пары", []int{1, 2, 3}, 100, 0},
		{"ноль как результат суммы", []int{1, -1, 0}, 0, 1},
		{"пустой вход", []int{}, 5, 0},
		{"цели нет", []int{5, 5, 5}, 10, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := TwoSumCount(tc.nums, tc.target); got != tc.want {
				t.Errorf("TwoSumCount(%v, %d) = %d, want %d",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestTwoSumCountMatchesBruteForce(t *testing.T) {
	// Перебор в лоб как эталон.
	cases := [][]int{
		{1, 1, 1},
		{1, 2, 3, 4},
		{0, 0, 0, 0},
		{4, -2, 4, 0, -2, 1},
	}
	for _, nums := range cases {
		for target := -6; target <= 8; target++ {
			want := 0
			for i := 0; i < len(nums); i++ {
				for j := i + 1; j < len(nums); j++ {
					if nums[i]+nums[j] == target {
						want++
					}
				}
			}
			if got := TwoSumCount(nums, target); got != want {
				t.Errorf("TwoSumCount(%v, %d) = %d, brute force = %d",
					nums, target, got, want)
			}
		}
	}
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"классический пример", []int{1, 1, 1}, 2, 2},
		{"две непересекающиеся", []int{1, 2, 3}, 3, 2},
		{"нули в хвосте", []int{1, -1, 0}, 0, 3},
		{"всё с нуля", []int{0, 0, 0}, 0, 6},
		{"нет подмассива", []int{1, 2, 3}, 100, 0},
		{"один элемент равен k", []int{5}, 5, 1},
		{"отрицательные", []int{-1, -1, 1}, 0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SubarraySum(tc.nums, tc.k); got != tc.want {
				t.Errorf("SubarraySum(%v, %d) = %d, want %d",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}
