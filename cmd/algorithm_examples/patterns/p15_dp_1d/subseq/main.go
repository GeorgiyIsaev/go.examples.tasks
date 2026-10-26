package main

import "fmt"

func main() {
	fmt.Println("=== 15. DP 1D / Subsequences (подпоследовательности) ===")
	fmt.Println()

	for _, nums := range [][]int{
		{10, 9, 2, 5, 3, 7},
		{10, 9, 2, 5, 3, 7, 101, 18},
		{0, 1, 0, 3, 2, 3},
		{7, 7, 7, 7},
		{},
	} {
		fmt.Printf("LIS(%v) = %d (быстро: %d)\n",
			nums, LengthOfLIS(nums), LengthOfLISFast(nums))
	}
	fmt.Println()

	fmt.Println("IsSubsequence(\"abc\", \"ahbgdc\") =", IsSubsequence("abc", "ahbgdc")) // true
	fmt.Println("IsSubsequence(\"axc\", \"ahbgdc\") =", IsSubsequence("axc", "ahbgdc")) // false
	fmt.Println("IsSubsequence(\"\", \"abc\")       =", IsSubsequence("", "abc"))       // true
}
