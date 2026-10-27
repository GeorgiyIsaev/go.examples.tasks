package main

import "testing"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{1, 4, 2},
		{3, 3, 0},
		{0, 0, 0},
		{1, 0, 1},
		{255, 0, 8},
	}
	for _, tc := range tests {
		if got := HammingDistance(tc.a, tc.b); got != tc.want {
			t.Errorf("HammingDistance(%d, %d) = %d, want %d",
				tc.a, tc.b, got, tc.want)
		}
	}
}

func TestReverseBits(t *testing.T) {
	// Классический пример из условия LC190.
	if got := ReverseBits(43261596); got != 964176192 {
		t.Errorf("ReverseBits(43261596) = %d, want 964176192", got)
	}
	// 1 (младший бит) -> старший бит 32-битного числа.
	if got := ReverseBits(1); got != 1<<31 {
		t.Errorf("ReverseBits(1) = %d, want %d", got, uint32(1<<31))
	}
	if got := ReverseBits(0); got != 0 {
		t.Errorf("ReverseBits(0) = %d, want 0", got)
	}
}

// TestReverseBitsTwiceIsOriginal — реверс обращаем.
func TestReverseBitsTwiceIsOriginal(t *testing.T) {
	for _, x := range []uint32{0, 1, 43261596, 0xFFFFFFFF, 12345} {
		if got := ReverseBits(ReverseBits(x)); got != x {
			t.Errorf("ReverseBits twice of %d = %d", x, got)
		}
	}
}

func TestGetSum(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{5, 7, 12},
		{-1, 1, 0},
		{0, 0, 0},
		{-3, -4, -7},
		{100, -100, 0},
		{-5, 3, -2},
	}
	for _, tc := range tests {
		if got := GetSum(tc.a, tc.b); got != tc.want {
			t.Errorf("GetSum(%d, %d) = %d, want %d",
				tc.a, tc.b, got, tc.want)
		}
	}
}

// TestGetSumMatchesPlus — сверка с обычным сложением.
func TestGetSumMatchesPlus(t *testing.T) {
	pairs := [][2]int{
		{1, 2}, {-1, -1}, {1024, -1024},
		{-999, 1234}, {0, 7}, {2147483647, -2147483648},
	}
	for _, p := range pairs {
		if got, want := GetSum(p[0], p[1]), p[0]+p[1]; got != want {
			t.Errorf("GetSum(%d, %d) = %d, want %d", p[0], p[1], got, want)
		}
	}
}
