package main

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1 — счастливое", 1, true},
		{"7 — счастливое", 7, true},
		{"10 — счастливое", 10, true},
		{"19 — счастливое", 19, true},
		{"2 — несчастливое", 2, false},
		{"3 — несчастливое", 3, false},
		{"4 — несчастливое", 4, false},
		{"100 — счастливое", 100, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsHappy(tc.n); got != tc.want {
				t.Errorf("IsHappy(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
