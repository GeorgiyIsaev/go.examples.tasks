package main

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"классический пример",
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"один список", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"пустые списки игнорируются",
			[][]int{{}, {1, 2}, {}},
			[]int{1, 2}},
		{"все пустые", [][]int{{}, {}}, []int{}},
		{"ни одного списка", nil, []int{}},
		{"списки разной длины",
			[][]int{{1}, {2, 3, 4, 5}, {6, 7}},
			[]int{1, 2, 3, 4, 5, 6, 7}},
		{"одинаковые значения",
			[][]int{{2, 2}, {2}},
			[]int{2, 2, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tc.lists))
			for i, vals := range tc.lists {
				lists[i] = BuildList(vals)
			}
			got := ListToSlice(MergeKLists(lists))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MergeKLists(%v) = %v, want %v", tc.lists, got, tc.want)
			}
		})
	}
}

func TestMergeKListsProperties(t *testing.T) {
	// Результат отсортирован и содержит все элементы по разу.
	lists := [][]int{{1, 4, 7}, {2, 5}, {3, 6, 9}, {0}}
	all := []int{}
	nodes := make([]*ListNode, len(lists))
	for i, vals := range lists {
		nodes[i] = BuildList(vals)
		all = append(all, vals...)
	}
	got := ListToSlice(MergeKLists(nodes))

	if len(got) != len(all) {
		t.Fatalf("result has %d elements, want %d", len(got), len(all))
	}
	for i := 1; i < len(got); i++ {
		if got[i-1] > got[i] {
			t.Fatalf("result is not sorted: %v", got)
		}
	}
	// Мультимножество: сортируем и сравниваем.
	sortInts(all)
	sortInts(got)
	if !reflect.DeepEqual(got, all) {
		t.Errorf("result %v is not a permutation of input %v", got, all)
	}
}

// sortInts — простая сортировка вставками для проверок.
func sortInts(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
