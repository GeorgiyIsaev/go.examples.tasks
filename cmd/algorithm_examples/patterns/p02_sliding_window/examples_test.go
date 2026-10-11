package main

import (
	"reflect"
	"testing"
)

// ===== Fixed Window =====

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

// ===== Variable Window =====

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"пример 1", "abcabcbb", 3},
		{"пример 2", "bbbbb", 1},
		{"пример 3", "pwwkew", 3},
		{"пустая", "", 0},
		{"один символ", "a", 1},
		{"все уникальные", "abcdef", 6},
		{"палиндромный", "abba", 2},
		{"повтор внутри", "dvdf", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tc.s); got != tc.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d",
					tc.s, got, tc.want)
			}
		})
	}
}

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"пример", []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{"все единицы", []int{1, 1, 1}, 0, 3},
		{"все нули, k=0", []int{0, 0, 0}, 0, 0},
		{"все нули, k=2", []int{0, 0, 0}, 2, 2},
		{"пустой", []int{}, 1, 0},
		{"k больше нулей", []int{0, 0, 0}, 10, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongestOnes(tc.nums, tc.k); got != tc.want {
				t.Errorf("LongestOnes(%v, %d) = %d, want %d",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

func TestLengthOfLongestSubstringKDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"пример", "eceba", 2, 3},
		{"k = 1", "aa", 1, 2},
		{"k = 0", "abc", 0, 0},
		{"пустая", "", 2, 0},
		{"все уникальны, k=2", "abc", 2, 2},
		{"один символ, k=2", "aaaa", 2, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringKDistinct(tc.s, tc.k); got != tc.want {
				t.Errorf("LengthOfLongestSubstringKDistinct(%q, %d) = %d, want %d",
					tc.s, tc.k, got, tc.want)
			}
		})
	}
}

// ===== Min Window =====

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"пример", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"нет решения", 100, []int{1, 2, 3}, 0},
		{"одно число", 5, []int{5}, 1},
		{"всё окно", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"пустой", 5, []int{}, 0},
		{"ровно target", 4, []int{1, 4, 4}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinSubArrayLen(tc.target, tc.nums); got != tc.want {
				t.Errorf("MinSubArrayLen(%d, %v) = %d, want %d",
					tc.target, tc.nums, got, tc.want)
			}
		})
	}
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"пример", "ADOBECODEBANC", "ABC", "BANC"},
		{"t длиннее s", "a", "aa", ""},
		{"символы не найдены", "abc", "xyz", ""},
		{"s == t", "abc", "abc", "abc"},
		{"пустая t", "abc", "", ""},
		{"дубли в t", "aa", "aa", "aa"},
		{"окно в середине", "cabwefgewcwaefgcf", "cae", "cwae"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinWindow(tc.s, tc.t); got != tc.want {
				t.Errorf("MinWindow(%q, %q) = %q, want %q",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}
