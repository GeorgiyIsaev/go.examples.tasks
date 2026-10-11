package main

import (
	"reflect"
	"testing"
)

// ===== Opposite Ends =====

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

// ===== Slow / Fast =====

func TestRemoveDuplicatesSorted(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"пустой", []int{}, 0},
		{"один", []int{1}, 1},
		{"все одинаковые", []int{2, 2, 2, 2}, 1},
		{"все уникальные", []int{1, 2, 3, 4}, 4},
		{"с дублями", []int{1, 1, 2, 3, 3, 3, 4}, 4},
		{"дубли в конце", []int{1, 2, 3, 3, 3}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			got := RemoveDuplicatesSorted(nums)
			if got != tc.want {
				t.Errorf("RemoveDuplicatesSorted(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"пустой", []int{}, nil},
		{"без нулей", []int{1, 2, 3}, []int{1, 2, 3}},
		{"только нули", []int{0, 0, 0}, []int{0, 0, 0}},
		{"чередование", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"нули в конце", []int{1, 2, 3, 0, 0}, []int{1, 2, 3, 0, 0}},
		{"нули в начале", []int{0, 0, 1, 2, 3}, []int{1, 2, 3, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			MoveZeroes(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Errorf("MoveZeroes(%v) = %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		val  int
		want int
	}{
		{"пустой", []int{}, 1, 0},
		{"нет совпадений", []int{1, 2, 3}, 4, 3},
		{"всё совпадает", []int{3, 3, 3}, 3, 0},
		{"частично", []int{3, 2, 2, 3}, 3, 2},
		{"в середине", []int{1, 2, 3, 4, 2}, 2, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			got := RemoveElement(nums, tc.val)
			if got != tc.want {
				t.Errorf("RemoveElement(%v, %d) = %d, want %d", tc.nums, tc.val, got, tc.want)
			}
		})
	}
}

// ===== Three Pointers =====

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"пример из условия", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"все нули", []int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{"нет решений", []int{1, 2, 3}, [][]int{}},
		{"пустой", []int{}, [][]int{}},
		{"меньше трёх", []int{1, 2}, [][]int{}},
		{"много дублей", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ThreeSum(tc.nums)
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestSortColors(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"пустой", []int{}, nil},
		{"один", []int{1}, []int{1}},
		{"уже отсортирован", []int{0, 1, 2}, []int{0, 1, 2}},
		{"обратный", []int{2, 1, 0}, []int{0, 1, 2}},
		{"смешанный", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"только нули", []int{0, 0, 0}, []int{0, 0, 0}},
		{"только двойки", []int{2, 2, 2}, []int{2, 2, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			nums := append([]int(nil), tc.nums...)
			SortColors(nums)
			if !reflect.DeepEqual(nums, tc.want) {
				t.Errorf("SortColors(%v) = %v, want %v", tc.nums, nums, tc.want)
			}
		})
	}
}
