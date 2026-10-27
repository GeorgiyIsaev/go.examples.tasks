package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"классический пример", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"с дублями", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"k = 1 (максимум)", []int{7, 1, 9}, 1, 9},
		{"k = длина", []int{4, 1, 3, 2}, 4, 1},
		{"отрицательные", []int{-1, -2, -3}, 2, -2},
		{"один элемент", []int{5}, 1, 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := KthLargest(tc.nums, tc.k); got != tc.want {
				t.Errorf("KthLargest(%v, %d) = %d, want %d",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

func TestKthLargestStream(t *testing.T) {
	// Пример из LeetCode 703.
	s := NewKthLargestStream(3, []int{4, 5, 8, 2})
	got := []int{
		s.Add(3),
		s.Add(5),
		s.Add(10),
		s.Add(9),
		s.Add(4),
	}
	want := []int{4, 5, 5, 8, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Add sequence = %v, want %v", got, want)
	}
}

func TestKthLargestStreamMatchesSort(t *testing.T) {
	// Проверка против эталона: полная сортировка на каждом шаге.
	k := 2
	s := NewKthLargestStream(k, nil)
	stream := []int{10, 3, 8, 1, 7, 3, 100, -5}
	nums := []int{}
	for _, v := range stream {
		nums = append(nums, v)
		got := s.Add(v)
		if len(nums) < k {
			continue // пока элементов меньше k — сравнение не определено
		}
		sorted := append([]int{}, nums...)
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
		want := sorted[k-1]
		if got != want {
			t.Errorf("Add(%d): got %d, want %d (nums=%v)", v, got, want, nums)
		}
	}
}

func TestKClosest(t *testing.T) {
	// Порядок не гарантирован: нормализуем.
	normalize := func(points [][]int) [][]int {
		out := make([][]int, len(points))
		for i, p := range points {
			out[i] = append([]int{}, p...)
		}
		sort.Slice(out, func(a, b int) bool {
			if out[a][0] != out[b][0] {
				return out[a][0] < out[b][0]
			}
			return out[a][1] < out[b][1]
		})
		return out
	}

	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{"классический пример", [][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{"три точки", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{-2, 4}, {3, 3}}},
		{"k = все", [][]int{{1, 0}, {0, 1}}, 2, [][]int{{0, 1}, {1, 0}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := KClosest(tc.points, tc.k)
			if !reflect.DeepEqual(normalize(got), normalize(tc.want)) {
				t.Errorf("KClosest(%v, %d) = %v, want %v",
					tc.points, tc.k, got, tc.want)
			}
		})
	}
}
