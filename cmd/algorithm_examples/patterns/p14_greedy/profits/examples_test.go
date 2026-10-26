package main

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"классический пример", []int{7, 1, 5, 3, 6, 4}, 5},
		{"всё падает", []int{7, 6, 4, 3, 1}, 0},
		{"два дня рост", []int{1, 2}, 1},
		{"один день", []int{5}, 0},
		{"пусто", []int{}, 0},
		{"одинаковые цены", []int{3, 3, 3}, 0},
		{"минимум в конце", []int{2, 4, 1}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxProfit(tc.prices); got != tc.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}

func TestMaxProfitMany(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"классический пример", []int{7, 1, 5, 3, 6, 4}, 7},
		{"всё падает", []int{7, 6, 4, 3, 1}, 0},
		{"монотонный рост", []int{1, 2, 3, 4, 5}, 4},
		{"один день", []int{5}, 0},
		{"пусто", []int{}, 0},
		{"чередование", []int{1, 2, 1, 2, 1}, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxProfitMany(tc.prices); got != tc.want {
				t.Errorf("MaxProfitMany(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}

// TestMaxProfitManyBeatsSingle — неограниченные сделки никогда не
// хуже одной.
func TestMaxProfitManyBeatsSingle(t *testing.T) {
	cases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 9, 2, 9, 1},
		{5, 4, 3, 2, 1},
		{2, 2, 2},
	}
	for _, prices := range cases {
		single := MaxProfit(prices)
		many := MaxProfitMany(prices)
		if many < single {
			t.Errorf("prices=%v: many=%d < single=%d", prices, many, single)
		}
	}
}

// TestMaxProfitManyMatchesDP — эталон: классическая DP-формулировка
// (best[i] — максимум, заканчиваясь на день i) через перебор точки
// покупки.
func TestMaxProfitManyMatchesDP(t *testing.T) {
	brute := func(prices []int) int {
		total := 0
		// сумма всех положительных скачков — эквивалент «покупай
		// каждый подъём», сверим с перебором «все пары покупок»
		// через склейку сделок: максимум = сумма подъёмов.
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				total += prices[i] - prices[i-1]
			}
		}
		return total
	}

	cases := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{3, 2, 6, 5, 0, 3},
		{2, 4, 1},
	}
	for _, prices := range cases {
		want := brute(prices)
		if got := MaxProfitMany(prices); got != want {
			t.Errorf("MaxProfitMany(%v) = %d, reference = %d", prices, got, want)
		}
	}
}
