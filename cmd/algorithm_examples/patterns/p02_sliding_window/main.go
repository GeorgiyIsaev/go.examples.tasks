package main

import "fmt"

func main() {
	fmt.Println("=== 02. Sliding Window ===")
	fmt.Println()

	fmt.Println("--- Fixed Window ---")
	fmt.Printf("MaxSumSubarrayOfSizeK([2,1,5,1,3,2], 3)  = %d\n",
		MaxSumSubarrayOfSizeK([]int{2, 1, 5, 1, 3, 2}, 3))
	fmt.Printf("FindAverages([1,3,2,6,-1,4,1,8,2], 5)   = %v\n",
		FindAverages([]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5))
	fmt.Printf("MaxVowels(\"abciiidef\", 3)                = %d\n",
		MaxVowels("abciiidef", 3))

	fmt.Println()
	fmt.Println("--- Variable Window (максимум) ---")
	fmt.Printf("LengthOfLongestSubstring(\"abcabcbb\")      = %d\n",
		LengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("LengthOfLongestSubstring(\"bbbbb\")         = %d\n",
		LengthOfLongestSubstring("bbbbb"))
	fmt.Printf("LengthOfLongestSubstring(\"pwwkew\")        = %d\n",
		LengthOfLongestSubstring("pwwkew"))
	fmt.Printf("LongestOnes([1,1,1,0,0,0,1,1,1,1,0], 2)  = %d\n",
		LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Printf("LengthOfLongestSubstringKDistinct(\"eceba\", 2) = %d\n",
		LengthOfLongestSubstringKDistinct("eceba", 2))

	fmt.Println()
	fmt.Println("--- Min Window ---")
	fmt.Printf("MinSubArrayLen(7, [2,3,1,2,4,3])         = %d\n",
		MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Printf("MinSubArrayLen(100, [1,2,3])             = %d\n",
		MinSubArrayLen(100, []int{1, 2, 3}))
	fmt.Printf("MinWindow(\"ADOBECODEBANC\", \"ABC\")        = %q\n",
		MinWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("MinWindow(\"a\", \"aa\")                     = %q\n",
		MinWindow("a", "aa"))
}
