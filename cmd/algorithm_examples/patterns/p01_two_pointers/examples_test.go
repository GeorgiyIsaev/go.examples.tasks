package main

import (
	"reflect"
	"sort"
	"testing"
)

// ============================================================
// Тесты для twoSumSorted (вариант A — два конца)
// ============================================================

func TestTwoSumSorted(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		wantI  int
		wantJ  int
		wantOK bool
	}{
		{"обычный случай", []int{1, 2, 3, 4, 6, 8, 11}, 10, 1, 5, true},
		{"пара в начале", []int{1, 2, 3, 4, 5}, 3, 0, 1, true},
		{"пара в конце", []int{1, 2, 3, 4, 5}, 9, 3, 4, true},
		{"нет пары", []int{1, 2, 3, 4, 5}, 100, -1, -1, false},
		{"пустой массив", []int{}, 5, -1, -1, false},
		{"один элемент", []int{5}, 10, -1, -1, false},
		{"отрицательные числа", []int{-3, -1, 0, 2, 5}, 2, 0, 4, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, j, ok := twoSumSorted(tt.nums, tt.target)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v", ok, tt.wantOK)
			}
			if !ok {
				return
			}
			if i != tt.wantI || j != tt.wantJ {
				t.Fatalf("got (%d,%d), want (%d,%d)", i, j, tt.wantI, tt.wantJ)
			}
			// Дополнительная проверка: сумма действительно равна target
			if tt.nums[i]+tt.nums[j] != tt.target {
				t.Fatalf("sum %d+%d != %d", tt.nums[i], tt.nums[j], tt.target)
			}
		})
	}
}

// ============================================================
// Тесты для removeDuplicates (вариант B — slow/fast)
// ============================================================

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int // ожидаемые первые wantLen элементов
	}{
		{"пустой", []int{}, 0, nil},
		{"один", []int{1}, 1, []int{1}},
		{"нет дубликатов", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"все дубликаты", []int{2, 2, 2, 2}, 1, []int{2}},
		{"смешанные", []int{1, 1, 2, 2, 3, 4, 4}, 4, []int{1, 2, 3, 4}},
		{"дубликаты в конце", []int{1, 2, 3, 3, 3}, 3, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...) // копия, чтобы не портить test case
			got := removeDuplicates(nums)
			if got != tt.wantLen {
				t.Fatalf("len = %d, want %d", got, tt.wantLen)
			}
			gotNums := nums[:got]
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Fatalf("got %v, want %v", gotNums, tt.wantNums)
			}
		})
	}
}

// ============================================================
// Тесты для reverse (вариант C — разворот)
// ============================================================

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"пустой", []int{}, nil},
		{"один", []int{1}, []int{1}},
		{"два", []int{1, 2}, []int{2, 1}},
		{"нечётная длина", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"чётная длина", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)
			reverse(nums)
			if !reflect.DeepEqual(nums, tt.want) {
				t.Fatalf("got %v, want %v", nums, tt.want)
			}
		})
	}
}

// ============================================================
// Тесты для threeSum (тройка с суммой 0)
// ============================================================

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"пустой", []int{}, nil},
		{"меньше 3", []int{1, 2}, nil},
		{"нет троек", []int{1, 2, 3}, nil},
		{"классика", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"все нули", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			// Сортируем для стабильного сравнения
			sortResults(got)
			sortResults(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func sortResults(res [][]int) {
	for _, r := range res {
		sort.Ints(r)
	}
	sort.Slice(res, func(i, j int) bool {
		for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
			if res[i][k] != res[j][k] {
				return res[i][k] < res[j][k]
			}
		}
		return len(res[i]) < len(res[j])
	})
}
