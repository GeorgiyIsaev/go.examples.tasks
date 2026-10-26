package main

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"a", "", 0},
		{"AGGTAB", "GXTXAYB", 4},
		{"bsbininm", "jmjkbkjkv", 1},
	}
	for _, tc := range tests {
		if got := LongestCommonSubsequence(tc.s, tc.t); got != tc.want {
			t.Errorf("LongestCommonSubsequence(%q, %q) = %d, want %d",
				tc.s, tc.t, got, tc.want)
		}
	}
}

func TestLCSIsSymmetric(t *testing.T) {
	// ЛКС симметрична: LCS(a,b) == LCS(b,a).
	pairs := [][2]string{
		{"abcde", "ace"},
		{"AGGTAB", "GXTXAYB"},
		{"xyz", "yxz"},
	}
	for _, p := range pairs {
		ab := LongestCommonSubsequence(p[0], p[1])
		ba := LongestCommonSubsequence(p[1], p[0])
		if ab != ba {
			t.Errorf("LCS(%q,%q)=%d != LCS(%q,%q)=%d",
				p[0], p[1], ab, p[1], p[0], ba)
		}
	}
}

func TestMinDistance(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"", "abc", 3},
		{"abc", "", 3},
		{"same", "same", 0},
		{"a", "b", 1},
		{"kitten", "sitting", 3},
	}
	for _, tc := range tests {
		if got := MinDistance(tc.s, tc.t); got != tc.want {
			t.Errorf("MinDistance(%q, %q) = %d, want %d",
				tc.s, tc.t, got, tc.want)
		}
	}
}

func TestMinDistanceIsSymmetric(t *testing.T) {
	pairs := [][2]string{
		{"horse", "ros"},
		{"ab", "ba"},
		{"abc", "def"},
	}
	for _, p := range pairs {
		ab := MinDistance(p[0], p[1])
		ba := MinDistance(p[1], p[0])
		if ab != ba {
			t.Errorf("MinDistance(%q,%q)=%d != swapped %d",
				p[0], p[1], ab, ba)
		}
	}
}
