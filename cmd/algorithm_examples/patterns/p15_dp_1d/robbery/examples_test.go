package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{1, 2, 3, 1}, 4},
		{"второй пример LC", []int{2, 7, 9, 3, 1}, 12},
		{"один дом", []int{5}, 5},
		{"два дома", []int{1, 2}, 2},
		{"пустая улица", []int{}, 0},
		{"все нули", []int{0, 0, 0}, 0},
		{"убывающие", []int{9, 1, 1, 1}, 10},
		{"одинаковые", []int{3, 3, 3, 3}, 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Rob(tc.nums); got != tc.want {
				t.Errorf("Rob(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestRobCircle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{2, 3, 2}, 3},
		{"второй пример LC", []int{1, 2, 3, 1}, 4},
		{"третий пример: первый и последний соседи", []int{1, 2, 3}, 3},
		{"один дом", []int{5}, 5},
		{"два дома", []int{2, 1}, 2},
		{"два дома по кругу", []int{1, 2}, 2},
		{"все нули", []int{0, 0, 0, 0}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := RobCircle(tc.nums); got != tc.want {
				t.Errorf("RobCircle(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// TestRobCircleNeverTakesBothEnds — проверка сути ограничения:
// ни одно решение не содержит сразу nums[0] и nums[n-1].
func TestRobCircleNeverTakesBothEnds(t *testing.T) {
	// Эталонный перебор: все непустые независимые подмножества
	// индексов, без соседних по кругу.
	brute := func(nums []int) int {
		n := len(nums)
		if n == 0 {
			return 0
		}
		best := 0
		for mask := 0; mask < (1 << n); mask++ {
			ok := true
			sum := 0
			for i := 0; i < n; i++ {
				if mask&(1<<i) != 0 {
					sum += nums[i]
				}
			}
			// явная проверка всех соседних пар (включая первый-последний)
			for i := 0; i < n; i++ {
				j := (i + 1) % n
				if mask&(1<<i) != 0 && mask&(1<<j) != 0 {
					ok = false
					break
				}
			}
			if ok && sum > best {
				best = sum
			}
		}
		return best
	}

	cases := [][]int{
		{2, 3, 2},
		{1, 2, 3, 1},
		{1, 2, 3},
		{5, 5, 5, 5},
		{8, 2, 4, 1, 9},
	}
	for _, nums := range cases {
		want := brute(nums)
		if got := RobCircle(nums); got != want {
			t.Errorf("RobCircle(%v) = %d, brute force = %d", nums, got, want)
		}
	}
}
