package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{"классический пример", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{"все одинаковые", [][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{"уже без пересечений", [][]int{{1, 2}, {2, 3}, {3, 4}}, 0},
		{"пусто", [][]int{}, 0},
		{"один отрезок", [][]int{{1, 5}}, 0},
		{"вложенные", [][]int{{1, 10}, {2, 3}, {4, 5}}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := EraseOverlapIntervals(tc.intervals); got != tc.want {
				t.Errorf("EraseOverlapIntervals(%v) = %d, want %d",
					tc.intervals, got, tc.want)
			}
		})
	}
}

// TestEraseOverlapMatchesBrute — эталон: перебор всех подмножеств
// индексов, извлечь максимум непересекающихся.
func TestEraseOverlapMatchesBrute(t *testing.T) {
	maxNonOverlapping := func(intervals [][]int) int {
		n := len(intervals)
		best := 0
		for mask := 0; mask < (1 << n); mask++ {
			picked := []int{}
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					picked = append(picked, i)
				}
			}
			sort.Ints(picked)
			ok := true
			for k := 1; k < len(picked); k++ {
				prev := intervals[picked[k-1]]
				cur := intervals[picked[k]]
				if cur[0] < prev[1] { // пересечение (касание ок)
					ok = false
					break
				}
			}
			if ok && len(picked) > best {
				best = len(picked)
			}
		}
		return best
	}

	cases := [][][]int{
		{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
		{{1, 4}, {2, 3}, {3, 5}},
		{{1, 2}, {1, 2}, {1, 2}},
		{{1, 10}, {2, 3}, {4, 5}},
	}
	for _, intervals := range cases {
		want := len(intervals) - maxNonOverlapping(intervals)
		if got := EraseOverlapIntervals(intervals); got != want {
			t.Errorf("EraseOverlapIntervals(%v) = %d, brute = %d",
				intervals, got, want)
		}
	}
}

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{"классический пример", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{"все отдельно", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{"цепочка касаний", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{"один шар", [][]int{{1, 2}}, 1},
		{"пусто", [][]int{}, 0},
		{"все в одном месте", [][]int{{1, 1}, {1, 1}}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindMinArrowShots(tc.points); got != tc.want {
				t.Errorf("FindMinArrowShots(%v) = %d, want %d",
					tc.points, got, tc.want)
			}
		})
	}
}

// TestIntervalsGreedyDoesNotMutateInput — сортировка не должна
// менять входной срез.
func TestIntervalsGreedyDoesNotMutateInput(t *testing.T) {
	input := [][]int{{10, 16}, {2, 8}, {1, 6}}
	snapshot := make([][]int, len(input))
	for i, v := range input {
		snapshot[i] = append([]int{}, v...)
	}

	EraseOverlapIntervals(input)
	FindMinArrowShots(input)

	if !reflect.DeepEqual(input, snapshot) {
		t.Errorf("input mutated: %v, want %v", input, snapshot)
	}
}
