package main

import "testing"

// makeCyclic строит список и, если cycleAt >= 0, замыкает последний
// узел на узел с индексом cycleAt.
func makeCyclic(vals []int, cycleAt int) *ListNode {
	head := BuildList(vals)
	if head != nil && cycleAt >= 0 {
		tail := findNodeByIndex(head, len(vals)-1)
		tail.Next = findNodeByIndex(head, cycleAt)
	}
	return head
}

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
			if got := HasCycle(makeCyclic(tc.vals, tc.cycleAt)); got != tc.want {
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
			got := DetectCycle(makeCyclic(tc.vals, tc.cycleAt))
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
			if got := CountCycleLength(makeCyclic(tc.vals, tc.cycleAt)); got != tc.want {
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
