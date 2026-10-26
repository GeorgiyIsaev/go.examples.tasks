package main

import "testing"

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
