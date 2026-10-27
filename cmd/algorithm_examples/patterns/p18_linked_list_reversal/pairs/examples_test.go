package main

import (
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"классический пример", []int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{"чётная длина", []int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{"один элемент", []int{1}, []int{1}},
		{"два элемента", []int{1, 2}, []int{2, 1}},
		{"пустой список", nil, []int{}},
		{"три элемента — хвост не тронут", []int{1, 2, 3}, []int{2, 1, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ListToSlice(SwapPairs(BuildList(tc.vals)))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SwapPairs(%v) = %v, want %v", tc.vals, got, tc.want)
			}
		})
	}
}

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		{"двойки", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"тройки", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"k = длина", []int{1, 2, 3, 4}, 4, []int{4, 3, 2, 1}},
		{"k=1 — ничего не меняется", []int{1, 2, 3}, 1, []int{1, 2, 3}},
		{"хвост короче k", []int{1, 2, 3}, 5, []int{1, 2, 3}},
		{"хвост короче k (частично)", []int{1, 2, 3, 4, 5, 6, 7}, 3,
			[]int{3, 2, 1, 6, 5, 4, 7}},
		{"k больше длины", []int{1, 2}, 3, []int{1, 2}},
		{"пустой список", nil, 2, []int{}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ListToSlice(ReverseKGroup(BuildList(tc.vals), tc.k))
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReverseKGroup(%v, %d) = %v, want %v",
					tc.vals, tc.k, got, tc.want)
			}
		})
	}
}

// TestReverseKGroupNoCycle — защита от зацикливания: обход обязан
// завершиться и пройти ровно n узлов.
func TestReverseKGroupNoCycle(t *testing.T) {
	for _, k := range []int{1, 2, 3, 4, 5} {
		head := BuildList([]int{1, 2, 3, 4, 5, 6})
		res := ReverseKGroup(head, k)
		count := 0
		for cur := res; cur != nil; cur = cur.Next {
			count++
			if count > 6 {
				t.Fatalf("k=%d: cycle detected, list never ends", k)
			}
		}
		if count != 6 {
			t.Errorf("k=%d: visited %d nodes, want 6", k, count)
		}
	}
}

func TestSwapPairsDoubleSwapIsOriginal(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	once := SwapPairs(BuildList(vals))
	twice := ListToSlice(SwapPairs(once))
	if !reflect.DeepEqual(twice, vals) {
		t.Errorf("SwapPairs twice = %v, want %v", twice, vals)
	}
}
