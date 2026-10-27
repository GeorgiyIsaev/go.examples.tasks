package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{4, 1, 2, 1, 2}, 4},
		{"один элемент", []int{1}, 1},
		{"ноль как одиночка", []int{3, 3, 0}, 0},
		{"отрицательные", []int{-1, -1, 5}, 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SingleNumber(tc.nums); got != tc.want {
				t.Errorf("SingleNumber(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{2, 2, 3, 2}, 3},
		{"пример с нулём", []int{0, 1, 0, 1, 0, 1, 99}, 99},
		{"одиночка последняя", []int{5, 5, 5, 7}, 7},
		{"отрицательные", []int{-2, -2, -2, -3}, -3},
		{"все по три плюс одиночка", []int{7, 7, 7, 5}, 5},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SingleNumberII(tc.nums); got != tc.want {
				t.Errorf("SingleNumberII(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestSingleNumberIII(t *testing.T) {
	asSorted := func(a []int) []int {
		out := append([]int{}, a...)
		sort.Ints(out)
		return out
	}
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"классический пример", []int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{"ноль и минус", []int{-1, 0}, []int{-1, 0}},
		{"крупные", []int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SingleNumberIII(tc.nums)
			if !reflect.DeepEqual(asSorted(got), asSorted(tc.want)) {
				t.Errorf("SingleNumberIII(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

// TestSingleNumberIIIWorksForAnyPair — любые две одиночки находятся.
func TestSingleNumberIIIWorksForAnyPair(t *testing.T) {
	// LC260: две РАЗНЫЕ одиночки, всё остальное — пары.
	pairs := [][2]int{{0, 5}, {1, 7}, {-5, 3}, {100, -100}, {1 << 20, 3}}
	for _, p := range pairs {
		nums := []int{p[0], p[1], 5, 5, 9, 9, 42, 42}
		got := asSortedCopy(SingleNumberIII(nums))
		want := []int{p[0], p[1]}
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("pair %v: got %v", p, got)
		}
	}
}

func asSortedCopy(a []int) []int {
	out := append([]int{}, a...)
	sort.Ints(out)
	return out
}
