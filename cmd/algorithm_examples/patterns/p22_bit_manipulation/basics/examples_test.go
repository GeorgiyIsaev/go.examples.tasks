package main

import (
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		n    uint32
		want int
	}{
		{11, 3},  // 1011
		{255, 8}, // 11111111
		{0, 0},
		{1, 1},
		{8, 1},           // 1000
		{4294967295, 32}, // все 32 бита
	}
	for _, tc := range tests {
		if got := HammingWeight(tc.n); got != tc.want {
			t.Errorf("HammingWeight(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}

// TestHammingWeightMatchesNaive — сверка с наивным подсчётом битов.
func TestHammingWeightMatchesNaive(t *testing.T) {
	naive := func(n uint32) int {
		count := 0
		for n != 0 {
			if n&1 == 1 {
				count++
			}
			n >>= 1
		}
		return count
	}
	for _, n := range []uint32{0, 1, 7, 1024, 12345, 987654321} {
		if got, want := HammingWeight(n), naive(n); got != want {
			t.Errorf("HammingWeight(%d) = %d, naive = %d", n, got, want)
		}
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{2, true},
		{16, true},
		{1024, true},
		{18, false},
		{0, false},
		{-2, false},
		{6, false},
	}
	for _, tc := range tests {
		if got := IsPowerOfTwo(tc.n); got != tc.want {
			t.Errorf("IsPowerOfTwo(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}

func TestCountBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		{0, []int{0}},
		{8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	for _, tc := range tests {
		got := CountBits(tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("CountBits(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}

// TestCountBitsMatchesNaive — сверка с наивным подсчётом.
func TestCountBitsMatchesNaive(t *testing.T) {
	naive := func(x int) int {
		count := 0
		for x > 0 {
			count += x & 1
			x >>= 1
		}
		return count
	}
	for n := 0; n <= 64; n++ {
		got := CountBits(n)
		for i := 0; i <= n; i++ {
			if got[i] != naive(i) {
				t.Fatalf("CountBits(%d)[%d] = %d, want %d", n, i, got[i], naive(i))
			}
		}
	}
}
