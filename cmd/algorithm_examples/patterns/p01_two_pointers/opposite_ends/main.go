package main

import "fmt"

// Запуск: go run ./opposite_ends
func main() {
	fmt.Println("=== 01. Two Pointers / Opposite Ends (навстречу) ===")
	fmt.Println()

	s := []byte("hello")
	ReverseString(s)
	fmt.Printf("ReverseString(\"hello\")                          = %q\n", string(s))

	fmt.Printf("IsPalindrome(\"A man, a plan, a canal: Panama\")   = %v\n",
		IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("IsPalindrome(\"race a car\")                       = %v\n",
		IsPalindrome("race a car"))

	fmt.Printf("TwoSumSorted([2,7,11,15], 9)                    = %v\n",
		TwoSumSorted([]int{2, 7, 11, 15}, 9))
	fmt.Printf("TwoSumSorted([2,7,11,15], 1)                    = %v\n",
		TwoSumSorted([]int{2, 7, 11, 15}, 1))
	fmt.Printf("TwoSumSorted([-3,-1,0,4,5], 2)                  = %v\n",
		TwoSumSorted([]int{-3, -1, 0, 4, 5}, 2))
}
