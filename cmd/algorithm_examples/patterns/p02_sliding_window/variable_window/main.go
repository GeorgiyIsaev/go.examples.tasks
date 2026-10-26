package main

import "fmt"

// Запуск: go run ./variable_window
func main() {
	fmt.Println("=== 02. Sliding Window / Variable (растущее окно, максимум) ===")
	fmt.Println()

	fmt.Printf("LengthOfLongestSubstring(\"abcabcbb\")           = %d\n",
		LengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("LengthOfLongestSubstring(\"bbbbb\")              = %d\n",
		LengthOfLongestSubstring("bbbbb"))
	fmt.Printf("LengthOfLongestSubstring(\"pwwkew\")             = %d\n",
		LengthOfLongestSubstring("pwwkew"))
	fmt.Printf("LengthOfLongestSubstring(\"abba\")               = %d\n",
		LengthOfLongestSubstring("abba"))

	fmt.Printf("LongestOnes([1,1,1,0,0,0,1,1,1,1,0], 2)       = %d\n",
		LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))

	fmt.Printf("LengthOfLongestSubstringKDistinct(\"eceba\", 2)  = %d\n",
		LengthOfLongestSubstringKDistinct("eceba", 2))
}
