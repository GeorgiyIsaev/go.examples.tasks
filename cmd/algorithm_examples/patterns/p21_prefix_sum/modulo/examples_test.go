package main

import "testing"

func TestSubarraySumDivByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"классический пример", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"не делится", []int{5}, 9, 0},
		{"один ноль", []int{0}, 3, 1},
		{"все кратны", []int{5, 10, 15}, 5, 6},
		{"отрицательные остатки", []int{-1, 2, 3}, 2, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := SubarraySumDivByK(tc.nums, tc.k); got != tc.want {
				t.Errorf("SubarraySumDivByK(%v, %d) = %d, want %d",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

// TestSubarraySumDivByKMatchesBrute — эталон: все подмассивы.
func TestSubarraySumDivByKMatchesBrute(t *testing.T) {
	brute := func(nums []int, k int) int {
		count := 0
		for i := range nums {
			sum := 0
			for j := i; j < len(nums); j++ {
				sum += nums[j]
				if sum%k == 0 {
					count++
				}
			}
		}
		return count
	}

	cases := []struct {
		nums []int
		k    int
	}{
		{[]int{4, 5, 0, -2, -3, 1}, 5},
		{[]int{-1, 2, 3}, 2},
		{[]int{0, 0, 0}, 3},
		{[]int{7, 3, -2, 6}, 4},
	}
	for _, c := range cases {
		want := brute(c.nums, c.k)
		if got := SubarraySumDivByK(c.nums, c.k); got != want {
			t.Errorf("SubarraySumDivByK(%v, %d) = %d, brute = %d",
				c.nums, c.k, got, want)
		}
	}
}

func TestFindMaxLength(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"два элемента", []int{0, 1}, 2},
		{"серединный подмассив", []int{0, 1, 0}, 2},
		{"почти весь", []int{0, 1, 1, 0, 1}, 4},
		{"нечего", []int{0}, 0},
		{"весь массив", []int{1, 1, 0, 0}, 4},
		{"пусто", []int{}, 0},
		{"без пар", []int{1, 1, 1}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := FindMaxLength(tc.nums); got != tc.want {
				t.Errorf("FindMaxLength(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// TestFindMaxLengthMatchesBrute — эталон: перебор всех подмассивов.
func TestFindMaxLengthMatchesBrute(t *testing.T) {
	brute := func(nums []int) int {
		best := 0
		for i := range nums {
			ones, zeros := 0, 0
			for j := i; j < len(nums); j++ {
				if nums[j] == 1 {
					ones++
				} else {
					zeros++
				}
				if ones == zeros && 2*ones > best {
					best = 2 * ones
				}
			}
		}
		return best
	}

	cases := [][]int{
		{0, 1, 0},
		{0, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 0},
		{0, 0, 1, 1, 1},
	}
	for _, nums := range cases {
		want := brute(nums)
		if got := FindMaxLength(nums); got != want {
			t.Errorf("FindMaxLength(%v) = %d, brute = %d", nums, got, want)
		}
	}
}
