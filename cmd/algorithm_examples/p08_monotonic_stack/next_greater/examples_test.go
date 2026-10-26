package main

import (
	"reflect"
	"testing"
)

func TestNextGreater(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"смешанный пример", []int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
		{"строго убывающий", []int{5, 4, 3, 2, 1}, []int{-1, -1, -1, -1, -1}},
		{"строго возрастающий", []int{1, 2, 3, 4}, []int{2, 3, 4, -1}},
		{"один элемент", []int{7}, []int{-1}},
		{"дубли: строго больший", []int{2, 2, 2}, []int{-1, -1, -1}},
		{"дубли через равный", []int{1, 1, 2}, []int{2, 2, -1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := nextGreater(tc.nums); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("nextGreater(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestNextGreaterElement(t *testing.T) {
	got := NextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	want := []int{-1, 3, -1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("NextGreaterElement(...) = %v, want %v", got, want)
	}

	if res := NextGreaterElement([]int{2}, []int{1, 2}); !reflect.DeepEqual(res, []int{-1}) {
		t.Errorf("NextGreaterElement single = %v, want [-1]", res)
	}
}

func TestNextGreaterElementsCircular(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"классический пример", []int{1, 2, 1}, []int{2, -1, 2}},
		{"возрастающий", []int{1, 2, 3, 4}, []int{2, 3, 4, -1}},
		{"убывающий — все видят 4 через конец", []int{4, 3, 2, 1}, []int{-1, 4, 4, 4}},
		{"один элемент", []int{5}, []int{-1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := NextGreaterElements(tc.nums); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("NextGreaterElements(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

// TestNextGreaterMatchesBruteForce — эталон: для каждого i ближайший
// больший линейным проходом.
func TestNextGreaterMatchesBruteForce(t *testing.T) {
	cases := [][]int{
		{2, 1, 2, 4, 3},
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4},
		{3, 3, 3},
		{7},
	}
	for _, nums := range cases {
		want := make([]int, len(nums))
		for i := range want {
			want[i] = -1
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > nums[i] {
					want[i] = nums[j]
					break
				}
			}
		}
		if got := nextGreater(nums); !reflect.DeepEqual(got, want) {
			t.Errorf("nextGreater(%v) = %v, brute force = %v", nums, got, want)
		}
	}
}
