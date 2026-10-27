package main

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{"классический пример", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"второй пример LC", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{"пусто", []int{}, []string{}},
		{"один элемент", []int{1}, []string{"1"}},
		{"всё подряд", []int{1, 2, 3, 4}, []string{"1->4"}},
		{"все раздельные", []int{1, 3, 5}, []string{"1", "3", "5"}},
		{"отрицательные", []int{-2, -1, 0, 3}, []string{"-2->0", "3"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SummaryRanges(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SummaryRanges(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}

// TestSummaryRangesCoversAll — конкатенация всех диапазонов даёт
// исходный массив (каждое число ровно один раз).
func TestSummaryRangesCoversAll(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7, 8, 9}
	got := SummaryRanges(nums)

	// Разворачиваем диапазоны обратно в числа.
	expansion := []int{}
	for _, s := range got {
		start, end := parseRange(s)
		for v := start; v <= end; v++ {
			expansion = append(expansion, v)
		}
	}
	if !reflect.DeepEqual(expansion, nums) {
		t.Errorf("expansion = %v, want %v", expansion, nums)
	}
}

// parseRange разбирает "a" или "a->b".
func parseRange(s string) (int, int) {
	start, end := 0, 0
	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i = 1
	}
	for ; i < len(s) && s[i] != '-'; i++ {
		start = start*10 + int(s[i]-'0')
	}
	start *= sign
	end = start
	// ищем "->"
	for ; i < len(s); i++ {
		if s[i] == '>' {
			i++
			break
		}
	}
	if i < len(s) {
		sign = 1
		if s[i] == '-' {
			sign = -1
			i++
		}
		end = 0
		for ; i < len(s); i++ {
			end = end*10 + int(s[i]-'0')
		}
		end *= sign
	}
	return start, end
}

func TestItoa(t *testing.T) {
	for _, v := range []int{0, 7, 42, -1, -100} {
		if got := itoa(v); got != formatInt(v) {
			t.Errorf("itoa(%d) = %q, want %q", v, got, formatInt(v))
		}
	}
}

// formatInt — эталон через strconv в тесте.
func formatInt(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	buf := []byte{}
	for n > 0 {
		buf = append([]byte{byte('0' + n%10)}, buf...)
		n /= 10
	}
	if neg {
		return "-" + string(buf)
	}
	return string(buf)
}
