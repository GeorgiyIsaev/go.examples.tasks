package main

import (
	"reflect"
	"testing"
)

func TestCharFrequency(t *testing.T) {
	got := CharFrequency("banana")
	want := map[rune]int{'b': 1, 'a': 3, 'n': 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CharFrequency(banana) = %v, want %v", got, want)
	}

	if empty := CharFrequency(""); len(empty) != 0 {
		t.Errorf("CharFrequency(\"\") = %v, want empty map", empty)
	}

	// Многобайтовые символы не должны резаться.
	utf := CharFrequency("привет")
	if utf['п'] != 1 || utf['и'] != 1 || len(utf) != 6 {
		t.Errorf("CharFrequency(привет) = %v, want 6 distinct runes", utf)
	}
}

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"классический пример", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"все одинаковые", []int{3, 3, 3}, 3},
		{"один элемент", []int{1}, 1},
		{"большинство в конце", []int{1, 2, 3, 3, 3}, 3},
		{"негативные", []int{-1, -1, -1, 2}, -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MajorityElement(tc.nums); got != tc.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	// Ответ неупорядочен: сверяем как множество.
	asSet := func(a []int) map[int]bool {
		m := map[int]bool{}
		for _, v := range a {
			m[v] = true
		}
		return m
	}

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"классический пример", []int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{"все уникальны", []int{4, 5, 6}, 2, []int{4, 5}}, // любые 2, проверим размер ниже
		{"один элемент", []int{7}, 1, []int{7}},
		{"k = все", []int{1, 1, 2, 2, 3}, 3, []int{1, 2, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopKFrequent(tc.nums, tc.k)
			if len(got) != tc.k {
				t.Fatalf("TopKFrequent(%v, %d) returned %d values, want %d",
					tc.nums, tc.k, len(got), tc.k)
			}
			if tc.name == "все уникальны" {
				return // при равных частотах любой ответ валиден
			}
			if !reflect.DeepEqual(asSet(got), asSet(tc.want)) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want set %v",
					tc.nums, tc.k, got, tc.want)
			}
		})
	}
}

func TestTopKFrequentOrdering(t *testing.T) {
	// Первый элемент обязан быть самым частым.
	got := TopKFrequent([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 2)
	want := []int{3, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TopKFrequent(...) = %v, want %v", got, want)
	}
}
