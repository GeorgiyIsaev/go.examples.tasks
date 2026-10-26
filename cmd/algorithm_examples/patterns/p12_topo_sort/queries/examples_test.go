package main

import (
	"reflect"
	"testing"
)

func TestCheckIfPrerequisite(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		prereq  [][]int
		queries [][]int
		want    []bool
	}{
		{
			"классический пример LC",
			4,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[][]int{{0, 1}, {0, 3}},
			[]bool{false, false},
		},
		{
			"цепочка 3->1->0",
			4,
			[][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[][]int{{3, 0}, {3, 1}, {1, 0}},
			[]bool{true, true, true},
		},
		{
			"без рёбер — себя не считаем предшественником",
			3,
			nil,
			[][]int{{0, 1}, {1, 1}},
			[]bool{false, false},
		},
		{
			"длинная цепочка",
			5,
			[][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			[][]int{{0, 4}, {4, 0}},
			[]bool{true, false},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckIfPrerequisite(tc.n, tc.prereq, tc.queries)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CheckIfPrerequisite() = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestBothApproachesAgree — перебор запросов и предвычислённая таблица
// обязаны отвечать одинаково.
func TestBothApproachesAgree(t *testing.T) {
	cases := []struct {
		n       int
		prereq  [][]int
		queries [][]int
	}{
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			[][]int{{0, 1}, {0, 3}, {3, 0}, {3, 1}, {2, 0}, {1, 2}}},
		{3, nil, [][]int{{0, 0}, {0, 1}, {2, 1}}},
		{5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			[][]int{{0, 4}, {4, 0}, {1, 3}, {3, 1}}},
	}

	for i, tc := range cases {
		byDFS := CheckIfPrerequisite(tc.n, tc.prereq, tc.queries)
		byTable := PredecessorQueries(tc.n, tc.prereq, tc.queries)
		if !reflect.DeepEqual(byDFS, byTable) {
			t.Errorf("case %d: DFS %v != table %v", i, byDFS, byTable)
		}
	}
}
