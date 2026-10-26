package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"простой пример", "()", true},
		{"смешанные", "([]{})", true},
		{"вложенные", "{[()()]}", true},
		{"порядок нарушен", "([)]", false},
		{"только открывающая", "(", false},
		{"только закрывающая", ")", false},
		{"чужая скобка", "]", false},
		{"лишний символ", "(a)", false},
		{"пустая строка", "", true},
		{"последняя не закрыта", "(((", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValid(tc.s); got != tc.want {
				t.Errorf("IsValid(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"классический пример", "3[a2[c]]", "accaccacc"},
		{"две скобки подряд", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"буквы до скобки", "abc3[cd]xyz", "abccdcdcdxyz"},
		{"многозначное число", "10[a]", "aaaaaaaaaa"},
		{"без скобок", "abc", "abc"},
		{"вложенность в 3 уровня", "2[a2[b2[c]]]", "abccbccabccbcc"},
		{"одна буква", "a", "a"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := DecodeString(tc.s); got != tc.want {
				t.Errorf("DecodeString(%q) = %q, want %q", tc.s, got, tc.want)
			}
		})
	}
}
