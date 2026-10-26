package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize приводит ответ к сравниваемому виду: сортирует элементы
// внутри подмножеств и сами подмножества — порядок в backtracking
// задачах не гарантирован.
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, sub := range result {
		c := append([]int{}, sub...)
		sort.Ints(c)
		out[i] = c
	}
	sort.Slice(out, func(a, b int) bool {
		for i := 0; i < len(out[a]) && i < len(out[b]); i++ {
			if out[a][i] != out[b][i] {
				return out[a][i] < out[b][i]
			}
		}
		return len(out[a]) < len(out[b])
	})
	return out
}

func TestSubsets(t *testing.T) {
	got := normalize(Subsets([]int{1, 2, 3}))
	want := normalize([][]int{
		{}, {1}, {2}, {3},
		{1, 2}, {1, 3}, {2, 3},
		{1, 2, 3},
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Subsets([1,2,3]) = %v, want %v", got, want)
	}

	if n := len(Subsets([]int{1, 2, 3, 4})); n != 16 {
		t.Errorf("len(Subsets 4 elems) = %d, want 16", n)
	}

	empty := normalize(Subsets([]int{}))
	if !reflect.DeepEqual(empty, [][]int{{}}) {
		t.Errorf("Subsets([]) = %v, want [[]]", empty)
	}
}

func TestSubsetsWithDup(t *testing.T) {
	got := normalize(SubsetsWithDup([]int{1, 2, 2}))
	want := normalize([][]int{
		{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2},
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SubsetsWithDup([1,2,2]) = %v, want %v", got, want)
	}

	// Все элементы одинаковые — подмножества только по длине.
	allDup := normalize(SubsetsWithDup([]int{2, 2, 2}))
	wantAllDup := normalize([][]int{{}, {2}, {2, 2}, {2, 2, 2}})
	if !reflect.DeepEqual(allDup, wantAllDup) {
		t.Errorf("SubsetsWithDup([2,2,2]) = %v, want %v", allDup, wantAllDup)
	}

	// Нет повторов между собой.
	res := SubsetsWithDup([]int{1, 2, 2})
	seen := map[string]bool{}
	for _, sub := range res {
		key := fmt.Sprint(sub)
		if seen[key] {
			t.Errorf("duplicate subset %v", sub)
		}
		seen[key] = true
	}
}
