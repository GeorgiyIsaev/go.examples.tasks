package main

import (
	"reflect"
	"testing"
)

func TestTraversals(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		pre  []int
		in   []int
		post []int
	}{
		{"полное дерево", []int{1, 2, 3, 4, 5, 6, 7},
			[]int{1, 2, 4, 5, 3, 6, 7},
			[]int{4, 2, 5, 1, 6, 3, 7},
			[]int{4, 5, 2, 6, 7, 3, 1}},
		{"один узел", []int{1},
			[]int{1},
			[]int{1},
			[]int{1}},
		{"пустое дерево", nil, []int{}, []int{}, []int{}},
		{"только левая ветка", []int{1, 2, -1, 3},
			[]int{1, 2, 3},
			[]int{3, 2, 1},
			[]int{3, 2, 1}},
		{"только правая ветка", []int{1, -1, 2, -1, 3},
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{3, 2, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := BuildTree(tc.vals)
			if got := PreOrder(root); !reflect.DeepEqual(got, tc.pre) {
				t.Errorf("PreOrder(%v) = %v, want %v", tc.vals, got, tc.pre)
			}
			if got := InOrder(root); !reflect.DeepEqual(got, tc.in) {
				t.Errorf("InOrder(%v) = %v, want %v", tc.vals, got, tc.in)
			}
			if got := PostOrder(root); !reflect.DeepEqual(got, tc.post) {
				t.Errorf("PostOrder(%v) = %v, want %v", tc.vals, got, tc.post)
			}
		})
	}
}

func TestIterativeMatchesRecursive(t *testing.T) {
	trees := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1},
		nil,
		{1, 2, -1, 3},
		{4, 2, 6, 1, 3, 5, 7},
		{1, -1, 2, -1, 3, -1, 4},
	}

	for _, vals := range trees {
		root := BuildTree(vals)
		if got, want := PreOrderIterative(root), PreOrder(root); !reflect.DeepEqual(got, want) {
			t.Errorf("PreOrderIterative(%v) = %v, want %v", vals, got, want)
		}
		if got, want := InOrderIterative(root), InOrder(root); !reflect.DeepEqual(got, want) {
			t.Errorf("InOrderIterative(%v) = %v, want %v", vals, got, want)
		}
	}
}

func TestInOrderBST(t *testing.T) {
	// InOrder в BST обязан давать возрастающую последовательность.
	bst := BuildTree([]int{4, 2, 6, 1, 3, 5, 7})
	got := InOrder(bst)
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("InOrder(BST) = %v, want %v", got, want)
	}
}
