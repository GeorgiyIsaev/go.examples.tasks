package main

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 2},
		{8, 2},
		{9, 3},
		{15, 3},
		{16, 4},
		{100, 10},
		{1_000_000, 1000},
	}
	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			if got := MySqrt(tc.x); got != tc.want {
				t.Errorf("MySqrt(%d) = %d, want %d", tc.x, got, tc.want)
			}
		})
	}
}

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"пример 1", []int{3, 6, 7, 11}, 8, 4},
		{"пример 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"пример 3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"одна куча", []int{10}, 1, 10},
		{"одна куча, много часов", []int{10}, 10, 1},
		{"все единицы", []int{1, 1, 1, 1}, 4, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinEatingSpeed(tc.piles, tc.h); got != tc.want {
				t.Errorf("MinEatingSpeed(%v, %d) = %d, want %d",
					tc.piles, tc.h, got, tc.want)
			}
		})
	}
}

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name    string
		weights []int
		days    int
		want    int
	}{
		{"пример", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 15},
		{"пример 2", []int{3, 2, 2, 4, 1, 4}, 3, 6},
		{"пример 3", []int{1, 2, 3, 1, 1}, 4, 3},
		{"один груз", []int{5}, 1, 5},
		{"много дней", []int{1, 2, 3}, 10, 3},
		{"один день", []int{1, 2, 3}, 1, 6},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := ShipWithinDays(tc.weights, tc.days); got != tc.want {
				t.Errorf("ShipWithinDays(%v, %d) = %d, want %d",
					tc.weights, tc.days, got, tc.want)
			}
		})
	}
}
