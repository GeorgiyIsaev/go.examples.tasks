package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize сортирует перестановки лексикографически для сравнения.
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, p := range result {
		out[i] = append([]int{}, p...)
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

// factorial для проверки количества перестановок.
func factorial(n int) int {
	r := 1
	for i := 2; i <= n; i++ {
		r *= i
	}
	return r
}

func TestPermute(t *testing.T) {
	got := normalize(Permute([]int{1, 2, 3}))
	want := normalize([][]int{
		{1, 2, 3}, {1, 3, 2},
		{2, 1, 3}, {2, 3, 1},
		{3, 1, 2}, {3, 2, 1},
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Permute([1,2,3]) = %v, want %v", got, want)
	}

	if n := len(Permute([]int{1, 2, 3, 4})); n != factorial(4) {
		t.Errorf("len(Permute 4 elems) = %d, want %d", n, factorial(4))
	}

	if res := Permute([]int{}); !reflect.DeepEqual(res, [][]int{{}}) {
		t.Errorf("Permute([]) = %v, want [[]]", res)
	}
}

func TestPermuteUnique(t *testing.T) {
	got := normalize(PermuteUnique([]int{1, 1, 2}))
	want := normalize([][]int{
		{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PermuteUnique([1,1,2]) = %v, want %v", got, want)
	}

	// Все одинаковые — ровно одна перестановка.
	allDup := PermuteUnique([]int{5, 5, 5})
	if !reflect.DeepEqual(allDup, [][]int{{5, 5, 5}}) {
		t.Errorf("PermuteUnique([5,5,5]) = %v, want [[5 5 5]]", allDup)
	}

	// Никаких дублей в ответе.
	res := PermuteUnique([]int{1, 2, 2})
	seen := map[string]bool{}
	for _, p := range res {
		key := fmt.Sprint(p)
		if seen[key] {
			t.Errorf("duplicate permutation %v", p)
		}
		seen[key] = true
	}
}
