package main

import (
	"reflect"
	"testing"
)

func TestTrieSearchAndStartsWith(t *testing.T) {
	tr := NewTrie()
	for _, w := range []string{"apple", "app", "apricot", "banana", "band"} {
		tr.Insert(w)
	}

	tests := []struct {
		name      string
		word      string
		wantSrch  bool
		wantPrefx bool
	}{
		{"точное слово", "apple", true, true},
		{"короткое слово отдельно", "app", true, true},
		{"только префикс", "appl", false, true},
		{"чужое слово", "apply", false, false},
		{"полное не слово", "bananaa", false, false},
		{"не a..z", "app1e", false, false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tr.Search(tc.word); got != tc.wantSrch {
				t.Errorf("Search(%q) = %v, want %v", tc.word, got, tc.wantSrch)
			}
			if got := tr.StartsWith(tc.word); got != tc.wantPrefx {
				t.Errorf("StartsWith(%q) = %v, want %v", tc.word, got, tc.wantPrefx)
			}
		})
	}
}

func TestTrieReinsertAndEmpty(t *testing.T) {
	tr := NewTrie()
	tr.Insert("go")
	tr.Insert("go") // повторная вставка
	if !tr.Search("go") {
		t.Error("после повторной вставки слово должно найтись")
	}

	empty := NewTrie()
	if empty.Search("x") || empty.StartsWith("x") {
		t.Error("пустой trie ничего не находит")
	}
	if got := empty.WordsWithPrefix("x"); len(got) != 0 {
		t.Errorf("WordsWithPrefix в пустом trie = %v, want []", got)
	}
}

func TestWordsWithPrefix(t *testing.T) {
	tr := NewTrie()
	for _, w := range []string{"apple", "app", "apricot", "banana", "band", "ap"} {
		tr.Insert(w)
	}

	tests := []struct {
		name   string
		prefix string
		want   []string
	}{
		{"два слова на префикс", "ap", []string{"ap", "app", "apple", "apricot"}},
		{"регистр не поддерживается (только a..z)", "ban", []string{"banana", "band"}},
		{"нет такого префикса", "z", nil},
		{"само слово является префиксом", "appl", []string{"apple"}},
		{"пустой префикс — всё словарь", "", []string{"ap", "app", "apple", "apricot", "banana", "band"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tr.WordsWithPrefix(tc.prefix); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("WordsWithPrefix(%q) = %v, want %v", tc.prefix, got, tc.want)
			}
		})
	}
}

func TestTrieMatchesMapSet(t *testing.T) {
	// Сверка с эталоном: обычный map[string]bool.
	words := []string{"a", "an", "ant", "antenna", "bee", "beetle", "cat", "dog", "do", "door"}
	tr := NewTrie()
	set := map[string]bool{}
	for _, w := range words {
		tr.Insert(w)
		set[w] = true
	}

	probes := []string{"a", "an", "ant", "anta", "be", "bee", "beet", "beetle", "d", "do", "doorway", "z", ""}
	for _, p := range probes {
		if got := tr.Search(p); got != set[p] {
			t.Errorf("Search(%q) = %v, map = %v", p, got, set[p])
		}
		wantPref := false
		for w := range set {
			if len(w) >= len(p) && w[:len(p)] == p {
				wantPref = true
				break
			}
		}
		if got := tr.StartsWith(p); got != wantPref {
			t.Errorf("StartsWith(%q) = %v, want %v", p, got, wantPref)
		}
	}
}
