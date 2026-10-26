package main

import "testing"

func TestKnapsack01(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{"классический пример", []int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7, 9},
		{"ничего не влезает", []int{5, 6}, []int{10, 10}, 4, 0},
		{"нулевая вместимость", []int{1, 3}, []int{1, 4}, 0, 0},
		{"один предмет", []int{3}, []int{5}, 3, 5},
		{"все влезают", []int{1, 2}, []int{3, 4}, 10, 7},
		{"выбор по ценности на вес", []int{2, 3, 4}, []int{3, 4, 5}, 5, 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Knapsack01(tc.weights, tc.values, tc.capacity)
			if got != tc.want {
				t.Errorf("Knapsack01(%v, %v, %d) = %d, want %d",
					tc.weights, tc.values, tc.capacity, got, tc.want)
			}
		})
	}
}

// TestKnapsackMatchesBruteForce — перебор всех масок предметов.
func TestKnapsackMatchesBruteForce(t *testing.T) {
	brute := func(weights, values []int, capacity int) int {
		n := len(weights)
		best := 0
		for mask := 0; mask < (1 << n); mask++ {
			w, v := 0, 0
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					w += weights[i]
					v += values[i]
				}
			}
			if w <= capacity && v > best {
				best = v
			}
		}
		return best
	}

	cases := []struct {
		w, v []int
		cap  int
	}{
		{[]int{1, 3, 4, 5}, []int{1, 4, 5, 7}, 7},
		{[]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 5},
		{[]int{1, 1, 1}, []int{1, 2, 3}, 2},
		{[]int{10}, []int{100}, 5},
	}
	for _, c := range cases {
		want := brute(c.w, c.v, c.cap)
		if got := Knapsack01(c.w, c.v, c.cap); got != want {
			t.Errorf("Knapsack01(%v, %v, %d) = %d, brute = %d",
				c.w, c.v, c.cap, got, want)
		}
	}
}

func TestSubsetSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"сумма набирается", []int{1, 5, 11, 5}, 11, true},
		{"не набирается", []int{1, 2, 3}, 7, false},
		{"ноль всегда", []int{1, 2}, 0, true},
		{"ровно один элемент", []int{3, 1, 4}, 4, true},
		{"больше суммы", []int{1, 2}, 10, false},
		{"ноль в массиве", []int{0, 1}, 1, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SubsetSum(tc.nums, tc.target); got != tc.want {
				t.Errorf("SubsetSum(%v, %d) = %v, want %v",
					tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestPartitionEqual(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"равные доли", []int{1, 5, 11, 5}, true},
		{"нечётная сумма", []int{1, 2, 3, 5}, false},
		{"два одинаковых", []int{2, 2}, true},
		{"три разных", []int{1, 2, 3}, true}, // 1+2 = 3
		{"нельзя", []int{1, 2, 5}, false},
		{"чётность суммы нечётная", []int{1}, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := PartitionEqual(tc.nums); got != tc.want {
				t.Errorf("PartitionEqual(%v) = %v, want %v",
					tc.nums, got, tc.want)
			}
		})
	}
}
