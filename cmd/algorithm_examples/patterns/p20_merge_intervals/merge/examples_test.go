package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{"классический пример", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"касание сливается", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"уже без пересечений", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"вложенные", [][]int{{1, 10}, {2, 3}}, [][]int{{1, 10}}},
		{"пусто", [][]int{}, [][]int{}},
		{"один", [][]int{{1, 1}}, [][]int{{1, 1}}},
		{"не по порядку", [][]int{{8, 10}, {1, 3}, {2, 6}}, [][]int{{1, 6}, {8, 10}}},
		{"все слились", [][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 4}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Merge(tc.intervals); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Merge(%v) = %v, want %v", tc.intervals, got, tc.want)
			}
		})
	}
}

func TestMergeDoesNotMutateInput(t *testing.T) {
	input := [][]int{{8, 10}, {1, 3}, {2, 6}}
	snapshot := [][]int{{8, 10}, {1, 3}, {2, 6}}
	Merge(input)
	if !reflect.DeepEqual(input, snapshot) {
		t.Errorf("input mutated: %v", input)
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		{"пересекает один", [][]int{{1, 3}, {6, 9}}, []int{2, 5},
			[][]int{{1, 5}, {6, 9}}},
		{"пересекает несколько", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{"вставляется в конец", [][]int{{1, 3}, {6, 9}}, []int{10, 12},
			[][]int{{1, 3}, {6, 9}, {10, 12}}},
		{"вставляется в начало", [][]int{{3, 5}, {7, 9}}, []int{1, 2},
			[][]int{{1, 2}, {3, 5}, {7, 9}}},
		{"всё сливается", [][]int{{1, 3}}, []int{2, 6}, [][]int{{1, 6}}},
		{"пустой список", nil, []int{5, 7}, [][]int{{5, 7}}},
		{"отдельно слева", [][]int{{5, 6}}, []int{1, 2}, [][]int{{1, 2}, {5, 6}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Insert(tc.intervals, tc.newInterval); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Insert(%v, %v) = %v, want %v",
					tc.intervals, tc.newInterval, got, tc.want)
			}
		})
	}
}

// TestInsertEqualsMergeOfAppend — Insert(x) должен совпадать с
// Merge(append(x, newInterval)): тот же результат двумя способами.
func TestInsertEqualsMergeOfAppend(t *testing.T) {
	cases := []struct {
		base  [][]int
		extra []int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}}, []int{4, 9}},
		{[][]int{{1, 4}}, []int{0, 0}},
		{[][]int{{2, 3}, {5, 7}}, []int{4, 6}},
	}
	for _, c := range cases {
		appended := make([][]int, 0, len(c.base)+1)
		appended = append(appended, c.base...)
		appended = append(appended, c.extra)

		viaInsert := Insert(c.base, c.extra)
		viaMerge := Merge(appended)
		if !reflect.DeepEqual(viaInsert, viaMerge) {
			t.Errorf("Insert(%v,%v) = %v, but Merge = %v",
				c.base, c.extra, viaInsert, viaMerge)
		}
	}
}
