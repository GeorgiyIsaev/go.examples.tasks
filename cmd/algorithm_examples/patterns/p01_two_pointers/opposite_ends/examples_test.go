package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"пустая строка", "", ""},
		{"один символ", "a", "a"},
		{"чётная длина", "abcd", "dcba"},
		{"нечётная длина", "hello", "olleh"},
		{"палиндром", "aba", "aba"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := []byte(tc.input)
			ReverseString(b)
			if got := string(b); got != tc.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"пустая", "", true},
		{"один символ", "a", true},
		{"простой палиндром", "aba", true},
		{"не палиндром", "abc", false},
		{"с пробелами и пунктуацией", "A man, a plan, a canal: Panama", true},
		{"не палиндром с пунктуацией", "race a car", false},
		{"только цифры палиндром", "12321", true},
		{"цифры не палиндром", "12345", false},
		{"только знаки", ".,;", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.input); got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"пример из условия", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"пара в конце", []int{1, 2, 3, 4, 6}, 10, []int{3, 4}},
		{"нет решения", []int{2, 7, 11, 15}, 1, nil},
		{"пустой массив", []int{}, 5, nil},
		{"один элемент", []int{1}, 1, nil},
		{"отрицательные", []int{-3, -1, 0, 4, 5}, 2, []int{0, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSumSorted(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TwoSumSorted(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
