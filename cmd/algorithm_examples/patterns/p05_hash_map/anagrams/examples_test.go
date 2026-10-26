package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"классический пример", "anagram", "nagaram", true},
		{"не анаграмма", "rat", "car", false},
		{"разная длина", "ab", "a", false},
		{"обе пустые", "", "", true},
		{"одна пустая", "", "a", false},
		{"с пробелом и регистром — считаем как есть", "a", "A", false},
		{"одинаковые", "abc", "abc", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAnagram(tc.s, tc.t); got != tc.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}

func TestIsAnagramUnicode(t *testing.T) {
	// Многобайтовые символы: «ёлка» и перестановка тех же букв.
	if !IsAnagram("лека", "елка") {
		t.Errorf("IsAnagram(лека, елка) = false, want true")
	}
	if IsAnagram("ёлка", "елка") {
		t.Errorf("IsAnagram(ёлка, елка) = true, want false (ё != е)")
	}
}

func TestGroupAnagrams(t *testing.T) {
	// Группы неупорядочены: нормализуем — сортируем строки в группе
	// и сами группы по их склеенному виду.
	normalize := func(groups [][]string) [][]string {
		keys := make([]string, len(groups))
		for i, g := range groups {
			cp := append([]string{}, g...)
			sort.Strings(cp)
			keys[i] = strings.Join(cp, ",")
		}
		sort.Strings(keys)
		out := make([][]string, len(keys))
		for i, k := range keys {
			out[i] = strings.Split(k, ",")
		}
		return out
	}

	got := normalize(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	want := normalize([][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GroupAnagrams = %v, want %v", got, want)
	}

	// Пустой вход — пустой ответ.
	if res := GroupAnagrams([]string{}); len(res) != 0 {
		t.Errorf("GroupAnagrams([]) = %v, want []", res)
	}

	// Одиночки без анаграмм — каждая в своей группе.
	single := GroupAnagrams([]string{"a", "b", "c"})
	if len(single) != 3 {
		t.Errorf("GroupAnagrams([a,b,c]) has %d groups, want 3", len(single))
	}
}

func TestGroupAnagramsPreservesAllWords(t *testing.T) {
	// Ни одно слово не должно потеряться или задвоиться.
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	count := map[string]int{}
	for _, g := range GroupAnagrams(words) {
		for _, w := range g {
			count[w]++
		}
	}
	for _, w := range words {
		if count[w] != 1 {
			t.Errorf("word %q appears %d times, want 1", w, count[w])
		}
	}
	if len(count) != len(words) {
		t.Errorf("grouped %d distinct words, want %d", len(count), len(words))
	}
}
