package main

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{10, 9, 2, 5, 3, 7}, 3},
		{"второй пример LC", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"третий пример LC", []int{0, 1, 0, 3, 2, 3}, 4},
		{"все одинаковые", []int{7, 7, 7, 7}, 1},
		{"возрастающие", []int{1, 2, 3, 4}, 4},
		{"убывающие", []int{4, 3, 2, 1}, 1},
		{"пустой массив", []int{}, 0},
		{"один элемент", []int{5}, 1},
		{"отрицательные", []int{-3, -2, -1, -4}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLIS(tc.nums); got != tc.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// TestLISFastMatchesQuadratic — O(n log n) версия обязана совпадать
// с O(n²) на всех кейсах (длина LIS одна и та же, сами
// подпоследовательности могут отличаться).
func TestLISFastMatchesQuadratic(t *testing.T) {
	cases := [][]int{
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 1, 4, 1, 5, 9, 2, 6, 5},
		{2, 2},
		{},
		{42},
	}
	for _, nums := range cases {
		slow := LengthOfLIS(nums)
		fast := LengthOfLISFast(nums)
		if slow != fast {
			t.Errorf("nums=%v: O(n^2) = %d, O(n log n) = %d", nums, slow, fast)
		}
	}
}

// TestLISFastAgainstBrute — перебор всех подпоследовательностей
// на маленьких массивах (2^n масок).
func TestLISFastAgainstBrute(t *testing.T) {
	brute := func(nums []int) int {
		n := len(nums)
		best := 0
		for mask := 0; mask < (1 << n); mask++ {
			last := int(-1 << 30)
			length := 0
			ok := true
			for i := 0; i < n; i++ {
				if mask&(1<<i) == 0 {
					continue
				}
				if nums[i] <= last {
					ok = false
					break
				}
				last = nums[i]
				length++
			}
			if ok && length > best {
				best = length
			}
		}
		return best
	}

	cases := [][]int{
		{1, 3, 2, 4},
		{4, 3, 2, 1},
		{2, 2, 1},
		{0, 1, 0, 3, 2, 3},
	}
	for _, nums := range cases {
		want := brute(nums)
		if got := LengthOfLISFast(nums); got != want {
			t.Errorf("LengthOfLISFast(%v) = %d, brute = %d", nums, got, want)
		}
	}
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"является", "abc", "ahbgdc", true},
		{"не является", "axc", "ahbgdc", false},
		{"пустая подстрока", "", "abc", true},
		{"пустая цель при непустой", "a", "", false},
		{"целиком совпадает", "abc", "abc", true},
		{"более длинная", "abcd", "abc", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSubsequence(tc.s, tc.t); got != tc.want {
				t.Errorf("IsSubsequence(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}
