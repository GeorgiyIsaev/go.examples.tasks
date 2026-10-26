package main

import "testing"

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
				t.Errorf("LongestOnes(%v, %d) = %d, want %d", tc.nums, tc.k, got, tc.want)
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
