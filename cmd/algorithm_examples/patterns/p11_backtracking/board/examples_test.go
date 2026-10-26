package main

import (
	"sort"
	"strings"
	"testing"
)

func TestCountNQueens(t *testing.T) {
	// Классические значения N-Queens.
	want := map[int]int{1: 1, 2: 0, 3: 0, 4: 2, 5: 10, 6: 4, 7: 40, 8: 92}
	for n, w := range want {
		if got := CountNQueens(n); got != w {
			t.Errorf("CountNQueens(%d) = %d, want %d", n, got, w)
		}
	}
}

func TestSolveNQueensValid(t *testing.T) {
	n := 5
	solutions := SolveNQueens(n)
	if len(solutions) != 10 {
		t.Fatalf("SolveNQueens(5) returned %d solutions, want 10", len(solutions))
	}

	seen := map[string]bool{}
	for _, s := range solutions {
		if len(s) != n {
			t.Fatalf("solution has %d rows, want %d", len(s), n)
		}
		cols := map[int]bool{}
		d1, d2 := map[int]bool{}, map[int]bool{}
		for r, row := range s {
			if len(row) != n {
				t.Fatalf("row %q has wrong length", row)
			}
			c := strings.IndexByte(row, 'Q')
			if c < 0 {
				t.Fatalf("row %q has no queen", row)
			}
			if strings.Count(row, "Q") != 1 {
				t.Fatalf("row %q must contain exactly one queen", row)
			}
			if cols[c] {
				t.Errorf("solution %v: two queens in column %d", s, c)
			}
			cols[c] = true
			if d1[r-c] {
				t.Errorf("solution %v: diagonal conflict r-c=%d", s, r-c)
			}
			d1[r-c] = true
			if d2[r+c] {
				t.Errorf("solution %v: diagonal conflict r+c=%d", s, r+c)
			}
			d2[r+c] = true
		}
		key := strings.Join(s, "/")
		if seen[key] {
			t.Errorf("duplicate solution %v", s)
		}
		seen[key] = true
	}
}

func TestExistWord(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	// board не должен мутироваться между кейсами: возвращаем копии.
	newBoard := func() [][]byte {
		out := make([][]byte, len(board))
		for i, row := range board {
			out[i] = append([]byte{}, row...)
		}
		return out
	}

	tests := []struct {
		name string
		word string
		want bool
	}{
		{"существует", "ABCCED", true},
		{"короткое", "SEE", true},
		{"в конце отказывает", "ABCB", false},
		{"пустое слово", "", false},
		{"буквы нет на доске", "XYZ", false},
		{"одна буква", "F", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ExistWord(newBoard(), tc.word)
			if got != tc.want {
				t.Errorf("ExistWord(%q) = %v, want %v", tc.word, got, tc.want)
			}
		})
	}
}

func TestExistWordDoesNotMutateBoard(t *testing.T) {
	board := [][]byte{
		{'A', 'B'},
		{'C', 'D'},
	}
	before := joinBoard(board)
	ExistWord(board, "ABCD")
	after := joinBoard(board)
	if before != after {
		t.Errorf("board mutated: %q -> %q", before, after)
	}
}

func joinBoard(b [][]byte) string {
	parts := make([]string, len(b))
	for i, row := range b {
		parts[i] = string(row)
	}
	sort.Strings(parts)
	return strings.Join(parts, "|")
}
