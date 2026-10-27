package main

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"чётная длина", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{"нечётная длина", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{"два элемента", []int{1, 2}, []int{1, 2}},
		{"один элемент", []int{1}, []int{1}},
		{"пустой список", nil, []int{}},
		{"шесть элементов", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.vals)
			ReorderList(head)
			got := ListToSlice(head)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReorderList(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

// TestReorderListPreservesMultiset — reorder меняет порядок, но не
// состав: элементы те же, ровно по разу.
func TestReorderListPreservesMultiset(t *testing.T) {
	vals := []int{1, 1, 2, 3, 5, 8, 13}
	head := BuildList(vals)
	ReorderList(head)
	got := ListToSlice(head)

	if len(got) != len(vals) {
		t.Fatalf("len = %d, want %d", len(got), len(vals))
	}
	sortIntsSimple(got)
	sorted := append([]int{}, vals...)
	sortIntsSimple(sorted)
	if !reflect.DeepEqual(got, sorted) {
		t.Errorf("elements changed: got %v, want %v", got, sorted)
	}
}

// TestReorderListNoCycle — обход обязан завершиться.
func TestReorderListNoCycle(t *testing.T) {
	head := BuildList([]int{1, 2, 3, 4, 5})
	ReorderList(head)
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
		if count > 5 {
			t.Fatal("cycle detected")
		}
	}
	if count != 5 {
		t.Errorf("visited %d nodes, want 5", count)
	}
}

// sortIntsSimple — сортировка вставками для эталонных проверок.
func sortIntsSimple(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
