package main

import (
	"reflect"
	"testing"
)

func TestIntervalIntersection(t *testing.T) {
	tests := []struct {
		name   string
		first  [][]int
		second [][]int
		want   [][]int
	}{
		{"классический пример",
			[][]int{{0, 2}, {5, 10}, {13, 15}, {16, 18}},
			[][]int{{1, 5}, {8, 12}, {15, 24}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 15}, {16, 18}}},
		{"нет пересечений",
			[][]int{{1, 3}}, [][]int{{5, 7}},
			[][]int{}},
		{"касание — точка",
			[][]int{{1, 5}}, [][]int{{5, 8}},
			[][]int{{5, 5}}},
		{"полное покрытие",
			[][]int{{1, 10}}, [][]int{{2, 3}, {4, 5}},
			[][]int{{2, 3}, {4, 5}}},
		{"один список пуст",
			[][]int{}, [][]int{{1, 2}},
			[][]int{}},
		{"оба пусты",
			[][]int{}, [][]int{},
			[][]int{}},
		{"вложенные",
			[][]int{{0, 10}}, [][]int{{2, 5}, {6, 8}},
			[][]int{{2, 5}, {6, 8}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IntervalIntersection(tc.first, tc.second)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("IntervalIntersection(%v, %v) = %v, want %v",
					tc.first, tc.second, got, tc.want)
			}
		})
	}
}

// TestIntersectionIsSymmetric — порядок списков не важен.
func TestIntersectionIsSymmetric(t *testing.T) {
	pairs := []struct{ a, b [][]int }{
		{[][]int{{0, 2}, {5, 10}}, [][]int{{1, 5}, {8, 12}}},
		{[][]int{{1, 3}}, [][]int{{2, 4}}},
	}
	for _, p := range pairs {
		ab := IntervalIntersection(p.a, p.b)
		ba := IntervalIntersection(p.b, p.a)
		if !reflect.DeepEqual(ab, ba) {
			t.Errorf("AB = %v, BA = %v", ab, ba)
		}
	}
}

// TestIntersectionPointsAreInsideBoth — каждая точка результата
// лежит в каком-то отрезке обоих входов.
func TestIntersectionPointsAreInsideBoth(t *testing.T) {
	first := [][]int{{0, 2}, {5, 10}, {13, 15}}
	second := [][]int{{1, 5}, {8, 12}}

	got := IntervalIntersection(first, second)
	inside := func(list [][]int, p int) bool {
		for _, iv := range list {
			if iv[0] <= p && p <= iv[1] {
				return true
			}
		}
		return false
	}
	for _, iv := range got {
		for p := iv[0]; p <= iv[1]; p++ {
			if !inside(first, p) || !inside(second, p) {
				t.Fatalf("point %d of %v not in both lists", p, iv)
			}
		}
	}
}
