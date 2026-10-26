package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize сортирует и элементы комбинаций, и сами комбинации.
func normalize(result [][]int) [][]int {
	out := make([][]int, len(result))
	for i, c := range result {
		cp := append([]int{}, c...)
		sort.Ints(cp)
		out[i] = cp
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

// binomial C(n, k) для проверки количества комбинаций.
func binomial(n, k int) int {
	if k > n-k {
		k = n - k
	}
	r := 1
	for i := 1; i <= k; i++ {
		r = r * (n - k + i) / i
	}
	return r
}

func TestCombine(t *testing.T) {
	got := normalize(Combine(4, 2))
	want := normalize([][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Combine(4,2) = %v, want %v", got, want)
	}

	if n := len(Combine(5, 3)); n != binomial(5, 3) {
		t.Errorf("len(Combine(5,3)) = %d, want %d", n, binomial(5, 3))
	}

	if res := Combine(3, 0); len(res) != 1 {
		t.Errorf("Combine(3,0) = %v, want one empty combination", res)
	}
}

func TestCombinationSum(t *testing.T) {
	got := normalize(CombinationSum([]int{2, 3, 6, 7}, 7))
	want := normalize([][]int{{2, 2, 3}, {7}})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CombinationSum([2,3,6,7], 7) = %v, want %v", got, want)
	}

	got2 := normalize(CombinationSum([]int{2, 3, 5}, 8))
	want2 := normalize([][]int{{2, 3, 3}, {2, 2, 2, 2}, {3, 5}})
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("CombinationSum([2,3,5], 8) = %v, want %v", got2, want2)
	}

	if res := CombinationSum([]int{2}, 1); len(res) != 0 {
		t.Errorf("CombinationSum([2], 1) = %v, want []", res)
	}
}

func TestCombinationSumII(t *testing.T) {
	got := normalize(CombinationSumII([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	want := normalize([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CombinationSumII(...) = %v, want %v", got, want)
	}

	// Никаких комбинаций с двумя одинаковыми «разными» 1,
	// хотя ввод содержит два единичных элемента.
	res := CombinationSumII([]int{1, 1, 2}, 3)
	seen := map[string]bool{}
	for _, c := range res {
		key := fmt.Sprint(c)
		if seen[key] {
			t.Errorf("duplicate combination %v", c)
		}
		seen[key] = true
	}
	if !reflect.DeepEqual(normalize(res), normalize([][]int{{1, 2}})) {
		t.Errorf("CombinationSumII([1,1,2], 3) = %v, want [[1 2]]", res)
	}
}
