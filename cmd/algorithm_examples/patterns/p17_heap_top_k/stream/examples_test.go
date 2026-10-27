package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestMedianFinder(t *testing.T) {
	m := NewMedianFinder()
	steps := []struct {
		val  int
		want float64
	}{
		{6, 6},
		{10, 8},
		{2, 6},
		{6, 6},
		{5, 6},
		{4, 5.5},
		{0, 5},
	}
	for i, s := range steps {
		m.AddNum(s.val)
		if got := m.FindMedian(); !almostEqual(got, s.want) {
			t.Errorf("step %d (after %d): FindMedian() = %v, want %v",
				i+1, s.val, got, s.want)
		}
	}
}

func TestMedianFinderAgainstSort(t *testing.T) {
	// Эталон: полная сортировка накопленного потока.
	vals := []int{5, 5, 5, 1, 9, -3, 0, 7, 2, 8}
	m := NewMedianFinder()
	seen := []int{}
	for _, v := range vals {
		m.AddNum(v)
		seen = append(seen, v)

		sorted := append([]int{}, seen...)
		sortIntsSimple(sorted)
		var want float64
		n := len(sorted)
		if n%2 == 1 {
			want = float64(sorted[n/2])
		} else {
			want = float64(sorted[n/2-1]+sorted[n/2]) / 2
		}
		if got := m.FindMedian(); !almostEqual(got, want) {
			t.Errorf("after %v: FindMedian() = %v, want %v", seen, got, want)
		}
	}
}

func TestMedianFinderSingleElement(t *testing.T) {
	m := NewMedianFinder()
	m.AddNum(42)
	if got := m.FindMedian(); !almostEqual(got, 42) {
		t.Errorf("FindMedian() = %v, want 42", got)
	}
}

func TestMovingAverage(t *testing.T) {
	ma := NewMovingAverage(3)
	steps := []struct {
		val  int
		want float64
	}{
		{1, 1},
		{3, 2},
		{5, 3},
		{7, 5},
		{9, 7},
	}
	for i, s := range steps {
		if got := ma.Next(s.val); !almostEqual(got, s.want) {
			t.Errorf("step %d (Next(%d)) = %v, want %v", i+1, s.val, got, s.want)
		}
	}
}

func TestMovingAverageWindowOne(t *testing.T) {
	ma := NewMovingAverage(1)
	if got := ma.Next(5); !almostEqual(got, 5) {
		t.Errorf("Next(5) = %v, want 5", got)
	}
	if got := ma.Next(8); !almostEqual(got, 8) {
		t.Errorf("Next(8) = %v, want 8", got)
	}
}

// sortIntsSimple — сортировка вставками для эталонных проверок.
func sortIntsSimple(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j] < a[j-1]; j-- {
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}
