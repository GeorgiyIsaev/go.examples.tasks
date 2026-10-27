package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"классический пример", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"два элемента", []int{1, 2}, []int{2, 1}},
		{"один элемент", []int{7}, []int{7}},
		{"пустой список", nil, []int{}},
		{"дубли", []int{2, 2, 2}, []int{2, 2, 2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ListToSlice(ReverseList(BuildList(tc.vals)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReverseList(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"классический пример", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"два элемента", []int{1, 2}, []int{2, 1}},
		{"один элемент", []int{7}, []int{7}},
		{"пустой список", nil, []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ListToSlice(ReverseListRecursive(BuildList(tc.vals)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReverseListRecursive(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

// TestReverseTwiceIsOriginal — разворот обращаем: A => rev => A.
func TestReverseTwiceIsOriginal(t *testing.T) {
	for _, vals := range [][]int{{1, 2, 3}, {1}, nil, {1, 2, 3, 4, 5, 6}} {
		once := ReverseList(BuildList(vals))
		twice := ListToSlice(ReverseList(once))
		// ListToSlice(nil) = []int{}, а vals может быть nil —
		// сравниваем с нормализованным эталоном.
		want := ListToSlice(BuildList(vals))
		if !reflect.DeepEqual(twice, want) {
			t.Errorf("reverse twice of %v = %v, want %v", vals, twice, want)
		}
	}
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name        string
		vals        []int
		left, right int
		want        []int
	}{
		{"классический пример", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"весь список", []int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
		{"один узел", []int{1, 2, 3, 4, 5}, 3, 3, []int{1, 2, 3, 4, 5}},
		{"хвост", []int{1, 2, 3}, 2, 3, []int{1, 3, 2}},
		{"голова", []int{1, 2, 3}, 1, 2, []int{2, 1, 3}},
		{"два элемента", []int{1, 2}, 1, 2, []int{2, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ListToSlice(ReverseBetween(BuildList(tc.vals), tc.left, tc.right))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReverseBetween(%v, %d, %d) = %v, want %v",
					tc.vals, tc.left, tc.right, got, tc.want)
			}
		})
	}
}

// TestReverseBetweenContainsSameElements — разворот не теряет узлы.
func TestReverseBetweenContainsSameElements(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5, 6}
	got := ListToSlice(ReverseBetween(BuildList(vals), 2, 5))
	if len(got) != len(vals) {
		t.Fatalf("result len = %d, want %d", len(got), len(vals))
	}
	// Неизменная часть до left и после right.
	if !reflect.DeepEqual(got[:1], []int{1}) || !reflect.DeepEqual(got[5:], []int{6}) {
		t.Errorf("untouched parts changed: %v", got)
	}
	// Середина — перевёрнутый кусок.
	if !reflect.DeepEqual(got[1:5], []int{5, 4, 3, 2}) {
		t.Errorf("middle part = %v, want [5 4 3 2]", got[1:5])
	}
}
