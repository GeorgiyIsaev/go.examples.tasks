package main

import "fmt"

func main() {
	fmt.Println("=== 16. DP 2D / LCS и расстояние Левенштейна ===")
	fmt.Println()

	pairs := [][2]string{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
		{"", "abc"},
	}
	for _, p := range pairs {
		fmt.Printf("LCS(%q, %q) = %d\n", p[0], p[1],
			LongestCommonSubsequence(p[0], p[1]))
	}
	fmt.Println()

	for _, p := range [][2]string{
		{"horse", "ros"},
		{"intention", "execution"},
		{"", "abc"},
		{"same", "same"},
	} {
		fmt.Printf("MinDistance(%q, %q) = %d\n", p[0], p[1],
			MinDistance(p[0], p[1]))
	}
}
