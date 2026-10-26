package main

import "testing"

func TestMinStack(t *testing.T) {
	st := NewMinStack()
	if st.Len() != 0 {
		t.Fatalf("new stack len = %d, want 0", st.Len())
	}

	st.Push(5)
	st.Push(3)
	st.Push(7)

	if got := st.Top(); got != 7 {
		t.Errorf("Top() = %d, want 7", got)
	}
	if got := st.GetMin(); got != 3 {
		t.Errorf("GetMin() = %d, want 3", got)
	}

	st.Push(1)
	if got := st.GetMin(); got != 1 {
		t.Errorf("after push 1: GetMin() = %d, want 1", got)
	}

	st.Pop() // снимаем 1
	if got := st.GetMin(); got != 3 {
		t.Errorf("after pop: GetMin() = %d, want 3 (минимум восстановился)", got)
	}

	st.Pop() // снимаем 7
	st.Pop() // снимаем 3
	if got := st.GetMin(); got != 5 {
		t.Errorf("almost empty: GetMin() = %d, want 5", got)
	}

	st.Pop()
	if st.Len() != 0 {
		t.Errorf("empty stack len = %d, want 0", st.Len())
	}
}

func TestMinStackDuplicates(t *testing.T) {
	st := NewMinStack()
	st.Push(2)
	st.Push(2)
	st.Push(2)
	if got := st.GetMin(); got != 2 {
		t.Errorf("GetMin() = %d, want 2", got)
	}
	st.Pop()
	st.Pop()
	if got := st.GetMin(); got != 2 {
		t.Errorf("after two pops GetMin() = %d, want 2 (дубль ещё в стеке)", got)
	}
	st.Pop()
	if st.Len() != 0 {
		t.Errorf("stack should be empty, len = %d", st.Len())
	}
}

func TestBackspaceCompare(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"классический пример", "ab#c", "ad#c", true},
		{"двойное стирание", "ab##", "a#b#", true},
		{"разные длины", "a#c", "b", false},
		{"обе пустые", "", "", true},
		{"стирание в начале", "#a", "a", true},
		{"стирать нечего", "a", "a", true},
		{"лишний хвост", "a", "b", false},
		{"только решётки", "###", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := BackspaceCompare(tc.s, tc.t); got != tc.want {
				t.Errorf("BackspaceCompare(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}

func TestBackspaceCompareMatchesBruteForce(t *testing.T) {
	// Алгоритм с конца как эталон (без стека).
	brute := func(s string) string {
		out := []rune{}
		skip := 0
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '#' {
				skip++
				continue
			}
			if skip > 0 {
				skip--
				continue
			}
			out = append(out, rune(s[i]))
		}
		// руны уже в обратном порядке — реверсим
		for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
			out[i], out[j] = out[j], out[i]
		}
		return string(out)
	}

	cases := [][2]string{
		{"ab#c", "ad#c"},
		{"a##b", "a#b#"},
		{"#", ""},
		{"x###y", "y"},
		{"abc#", "abc"},
	}
	for _, c := range cases {
		want := brute(c[0]) == brute(c[1])
		if got := BackspaceCompare(c[0], c[1]); got != want {
			t.Errorf("BackspaceCompare(%q, %q) = %v, brute force = %v",
				c[0], c[1], got, want)
		}
	}
}
