package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"классический пример", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"ноль в середине", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{"один элемент", []int{2}, []int{1}},
		{"два элемента", []int{3, 5}, []int{5, 3}},
		{"два нуля", []int{0, 0}, []int{0, 0}},
		{"отрицательные", []int{-2, -3}, []int{-3, -2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ProductExceptSelf(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ProductExceptSelf(%v) = %v, want %v",
					tc.nums, got, tc.want)
			}
		})
	}
}

// TestProductExceptSelfMatchesBrute — эталон: умножение всех, кроме i.
func TestProductExceptSelfMatchesBrute(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 4},
		{-1, 1, 0, -3, 3},
		{2, 0, 5},
		{-2, -3, -4},
	}
	for _, nums := range cases {
		want := make([]int, len(nums))
		for i := range nums {
			p := 1
			for j, v := range nums {
				if j != i {
					p *= v
				}
			}
			want[i] = p
		}
		if got := ProductExceptSelf(nums); !reflect.DeepEqual(got, want) {
			t.Errorf("ProductExceptSelf(%v) = %v, brute = %v", nums, got, want)
		}
	}
}
