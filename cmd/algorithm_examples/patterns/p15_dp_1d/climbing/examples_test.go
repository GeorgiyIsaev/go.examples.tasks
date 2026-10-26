package main

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 8},
		{6, 13},
		{10, 89},
		{45, 1_836_311_903},
	}
	for _, tc := range tests {
		if got := ClimbStairs(tc.n); got != tc.want {
			t.Errorf("ClimbStairs(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}

// TestClimbStairsIsFibonacci — сверка с эталонным рекурсивным
// подсчётом (малые n) и с законом Фибоначчи (большие).
func TestClimbStairsIsFibonacci(t *testing.T) {
	var ways func(n int) int
	ways = func(n int) int {
		if n <= 2 {
			return n
		}
		return ways(n-1) + ways(n-2)
	}
	for n := 1; n <= 20; n++ {
		if got, want := ClimbStairs(n), ways(n); got != want {
			t.Errorf("ClimbStairs(%d) = %d, recursive = %d", n, got, want)
		}
	}
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{"классический пример", []int{10, 15, 20}, 15},
		{"второй пример LC", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"две ступеньки дёшево", []int{1, 100}, 1},
		{"одна ступенька", []int{5}, 0},
		{"одинаковые", []int{2, 2, 2, 2}, 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinCostClimbingStairs(tc.cost); got != tc.want {
				t.Errorf("MinCostClimbingStairs(%v) = %d, want %d",
					tc.cost, got, tc.want)
			}
		})
	}
}

// TestMinCostMatchesBruteForce — рекурсивный перебор всех путей
// на малых массивах: dp считает то же самое, но с мемоизацией.
func TestMinCostMatchesBruteForce(t *testing.T) {
	cases := [][]int{
		{10, 15, 20},
		{1, 100, 1, 1},
		{0, 0, 0},
		{5, 1, 2, 3},
	}
	for _, cost := range cases {
		n := len(cost)
		var best func(i int) int // стоимость пути из i до вершины
		best = func(i int) int {
			if i >= n {
				return 0 // вершина достигнута, платить не надо
			}
			one := best(i + 1)
			two := best(i + 2)
			if one < two {
				return cost[i] + one
			}
			return cost[i] + two
		}
		// Старт можно поставить на 0-ю или 1-ю ступеньку (бесплатно).
		want := best(0)
		if n > 1 {
			if b := best(1); b < want {
				want = b
			}
		}
		if got := MinCostClimbingStairs(cost); got != want {
			t.Errorf("MinCostClimbingStairs(%v) = %d, brute force = %d",
				cost, got, want)
		}
	}
}
