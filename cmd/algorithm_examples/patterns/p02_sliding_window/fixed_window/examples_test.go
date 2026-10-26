package main

import (
	"reflect"
	"testing"
)

func TestMaxSumSubarrayOfSizeK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"пример", []int{2, 1, 5, 1, 3, 2}, 3, 9},
		{"одно окно", []int{1, 2, 3}, 3, 6},
		{"пустой", []int{}, 3, 0},
		{"k больше массива", []int{1, 2}, 5, 0},
		{"k = 0", []int{1, 2, 3}, 0, 0},
		{"k = 1", []int{1, 5, 3}, 1, 5},
		{"отрицательные", []int{-1, -2, -3, -4}, 2, -3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxSumSubarrayOfSizeK(tc.nums, tc.k); got != tc.want {
				t.Errorf("MaxSumSubarrayOfSizeK(%v, %d) = %d, want %d",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

func TestFindAverages(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []float64
	}{
		{"пример", []int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5,
			[]float64{2.2, 2.8, 2.4, 3.6, 2.8}},
		{"одно окно", []int{1, 2, 3}, 3, []float64{2.0}},
		{"k = 1", []int{1, 2}, 1, []float64{1.0, 2.0}},
		{"пустой", []int{}, 3, []float64{}},
		{"k больше", []int{1, 2}, 5, []float64{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FindAverages(tc.nums, tc.k)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FindAverages(%v, %d) = %v, want %v",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"пример", "abciiidef", 3, 3},
		{"все гласные", "aeiou", 3, 3},
		{"нет гласных", "bcdfg", 2, 0},
		{"k = 1, есть гласная", "abc", 1, 1},
		{"k = 1, нет гласных", "bcd", 1, 0},
		{"пустая", "", 2, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxVowels(tc.s, tc.k); got != tc.want {
				t.Errorf("MaxVowels(%q, %d) = %d, want %d", tc.s, tc.k, got, tc.want)
			}
		})
	}
}
