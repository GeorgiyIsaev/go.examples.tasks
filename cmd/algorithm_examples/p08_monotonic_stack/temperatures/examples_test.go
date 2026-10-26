package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		temps []int
		want  []int
	}{
		{"классический пример", []int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"всё холодает", []int{9, 8, 7}, []int{0, 0, 0}},
		{"один день", []int{5}, []int{0}},
		{"возрастающие", []int{1, 2, 3, 4}, []int{1, 1, 1, 0}},
		{"дубли не считается warmer", []int{30, 30, 30}, []int{0, 0, 0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := DailyTemperatures(tc.temps); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("DailyTemperatures(%v) = %v, want %v", tc.temps, got, tc.want)
			}
		})
	}
}

// TestDailyTemperaturesBruteForce — эталонный линейный поиск.
func TestDailyTemperaturesBruteForce(t *testing.T) {
	cases := [][]int{
		{73, 74, 75, 71, 69, 72, 76, 73},
		{30, 40, 50, 60},
		{60, 50, 40, 30},
		{1, 1, 1, 1},
	}
	for _, temps := range cases {
		want := make([]int, len(temps))
		for i := range want {
			want[i] = 0
			for j := i + 1; j < len(temps); j++ {
				if temps[j] > temps[i] {
					want[i] = j - i
					break
				}
			}
		}
		if got := DailyTemperatures(temps); !reflect.DeepEqual(got, want) {
			t.Errorf("DailyTemperatures(%v) = %v, brute force = %v", temps, got, want)
		}
	}
}

func TestWarmDays(t *testing.T) {
	if got := WarmDays([]int{73, 74, 75, 71, 69, 72, 76, 73}); got != 10 {
		t.Errorf("WarmDays = %d, want 10 (1+1+4+2+1+1)", got)
	}
	if got := WarmDays([]int{9, 8, 7}); got != 0 {
		t.Errorf("WarmDays(cold) = %d, want 0", got)
	}
}
