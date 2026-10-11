package main

import (
	"testing"
)

// ===== Cycle Detection =====

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		cycleAt int // индекс узла, на который указывает последний; -1 = нет цикла
		want    bool
	}{
		{"пустой", nil, -1, false},
		{"один без цикла", []int{1}, -1, false},
		{"один с самоссылкой", []int{1}, 0, true},
		{"без цикла", []int{1, 2, 3, 4}, -1, false},
		{"цикл в середине", []int{1, 2, 3, 4}, 1, true},
		{"цикл в начале", []int{1, 2, 3, 4}, 0, true},
		{"цикл в конце", []int{1, 2, 3, 4}, 3, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			if head != nil && tc.cycleAt >= 0 {
				tail := findNodeByIndex(head, len(tc.vals)-1)
				tail.Next = findNodeByIndex(head, tc.cycleAt)
			}
			if got := HasCycle(head); got != tc.want {
				t.Errorf("HasCycle = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		cycleAt int
		wantVal int // -1 = nil
	}{
		{"пустой", nil, -1, -1},
		{"без цикла", []int{1, 2, 3, 4}, -1, -1},
		{"цикл на 2", []int{3, 2, 0, -4}, 1, 2},
		{"цикл на начало", []int{1, 2, 3, 4}, 0, 1},
		{"самоссылка", []int{1}, 0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			if head != nil && tc.cycleAt >= 0 {
				tail := findNodeByIndex(head, len(tc.vals)-1)
				tail.Next = findNodeByIndex(head, tc.cycleAt)
			}
			got := DetectCycle(head)
			if tc.wantVal == -1 {
				if got != nil {
					t.Errorf("DetectCycle = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Errorf("DetectCycle = nil, want node with Val=%d", tc.wantVal)
				return
			}
			if got.Val != tc.wantVal {
				t.Errorf("DetectCycle.Val = %d, want %d", got.Val, tc.wantVal)
			}
		})
	}
}

func TestCountCycleLength(t *testing.T) {
	tests := []struct {
		name    string
		vals    []int
		cycleAt int
		want    int
	}{
		{"пустой", nil, -1, 0},
		{"без цикла", []int{1, 2, 3, 4}, -1, 0},
		{"самоссылка", []int{1}, 0, 1},
		{"цикл длины 3", []int{3, 2, 0, -4}, 1, 3},
		{"цикл длины 4", []int{1, 2, 3, 4}, 0, 4},
		{"цикл длины 2", []int{1, 2, 3, 4}, 2, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			if head != nil && tc.cycleAt >= 0 {
				tail := findNodeByIndex(head, len(tc.vals)-1)
				tail.Next = findNodeByIndex(head, tc.cycleAt)
			}
			if got := CountCycleLength(head); got != tc.want {
				t.Errorf("CountCycleLength = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"пример 1", []int{1, 3, 4, 2, 2}, 2},
		{"пример 2", []int{3, 1, 3, 4, 2}, 3},
		{"минимальный", []int{1, 1}, 1},
		{"дубль в конце", []int{1, 2, 3, 4, 4}, 4},
		{"дубль в начале", []int{2, 2, 3, 4, 1}, 2},
		{"длинный", []int{1, 4, 3, 5, 6, 2, 4}, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindDuplicate(tc.nums); got != tc.want {
				t.Errorf("FindDuplicate(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// ===== Middle & Palindrome =====

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

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1 — счастливое", 1, true},
		{"7 — счастливое", 7, true},
		{"10 — счастливое", 10, true},
		{"19 — счастливое", 19, true},
		{"2 — несчастливое", 2, false},
		{"3 — несчастливое", 3, false},
		{"4 — несчастливое", 4, false},
		{"100 — счастливое", 100, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsHappy(tc.n); got != tc.want {
				t.Errorf("IsHappy(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
