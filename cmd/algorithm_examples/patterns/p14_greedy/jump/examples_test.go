package main

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"классический пример", []int{2, 3, 1, 1, 4}, true},
		{"дыра в конце", []int{3, 2, 1, 0, 4}, false},
		{"один элемент", []int{0}, true},
		{"ноль в начале", []int{0, 1}, false},
		{"все нули кроме последнего", []int{1, 0, 0}, false},
		{"один большой прыжок", []int{5, 0, 0, 0, 0}, true},
		{"ровно по краю", []int{1, 1, 1, 1}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanJump(tc.nums); got != tc.want {
				t.Errorf("CanJump(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

func TestJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{2, 3, 1, 1, 4}, 2},
		{"второй пример LC", []int{2, 3, 0, 1, 4}, 2},
		{"один элемент", []int{1}, 0},
		{"два элемента", []int{1, 1}, 1},
		{"всё сразу", []int{4, 1, 1, 1, 1}, 1},
		{"по одному", []int{1, 1, 1, 1}, 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Jump(tc.nums); got != tc.want {
				t.Errorf("Jump(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// TestJumpMatchesBFS — эталон: BFS по «уровням прыжков».
func TestJumpMatchesBFS(t *testing.T) {
	bfs := func(nums []int) int {
		n := len(nums)
		visited := make([]bool, n)
		type node struct{ idx, dist int }
		queue := []node{{0, 0}}
		visited[0] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.idx == n-1 {
				return cur.dist
			}
			for step := 1; step <= nums[cur.idx]; step++ {
				nb := cur.idx + step
				if nb >= n || visited[nb] {
					continue
				}
				visited[nb] = true
				queue = append(queue, node{nb, cur.dist + 1})
			}
		}
		return -1
	}

	cases := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
		{1, 1, 1, 1, 1},
		{5, 1, 1, 1, 1},
	}
	for _, nums := range cases {
		want := bfs(nums)
		if got := Jump(nums); got != want {
			t.Errorf("Jump(%v) = %d, BFS = %d", nums, got, want)
		}
	}
}
