package main

import "testing"

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"один", []int{1}, 1},
		{"два", []int{1, 2}, 2},
		{"три", []int{1, 2, 3}, 2},
		{"четыре", []int{1, 2, 3, 4}, 3},
		{"пять", []int{1, 2, 3, 4, 5}, 3},
		{"шесть", []int{1, 2, 3, 4, 5, 6}, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			if got := MiddleNode(head).Val; got != tc.want {
				t.Errorf("MiddleNode(%v).Val = %d, want %d", tc.vals, got, tc.want)
			}
		})
	}
}

func TestIsPalindromeList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"nil", nil, true},
		{"один", []int{1}, true},
		{"два одинаковых", []int{1, 1}, true},
		{"два разных", []int{1, 2}, false},
		{"чётный палиндром", []int{1, 2, 2, 1}, true},
		{"нечётный палиндром", []int{1, 2, 3, 2, 1}, true},
		{"не палиндром", []int{1, 2, 3, 4}, false},
		{"все одинаковые", []int{7, 7, 7, 7}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			if got := IsPalindromeList(head); got != tc.want {
				t.Errorf("IsPalindromeList(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}
