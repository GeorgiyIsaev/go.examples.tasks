package main

import (
	"reflect"
	"testing"
)

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
