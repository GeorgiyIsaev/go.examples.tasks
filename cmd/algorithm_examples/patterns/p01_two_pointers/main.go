package main

import "fmt"

func main() {
	fmt.Println("=== 01. Two Pointers ===")
	fmt.Println()

	fmt.Println("--- Opposite Ends ---")
	s := []byte("hello")
	ReverseString(s)
	fmt.Printf("ReverseString(\"hello\")                        = %q\n", string(s))
	fmt.Printf("IsPalindrome(\"A man, a plan, a canal: Panama\") = %v\n",
		IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("IsPalindrome(\"race a car\")                     = %v\n", IsPalindrome("race a car"))
	fmt.Printf("TwoSumSorted([2,7,11,15], 9)                  = %v\n", TwoSumSorted([]int{2, 7, 11, 15}, 9))
	fmt.Printf("TwoSumSorted([2,7,11,15], 1)                  = %v\n", TwoSumSorted([]int{2, 7, 11, 15}, 1))

	fmt.Println()
	fmt.Println("--- Slow / Fast ---")
	nums1 := []int{1, 1, 2, 3, 3, 3, 4}
	k := RemoveDuplicatesSorted(nums1)
	fmt.Printf("RemoveDuplicatesSorted([1,1,2,3,3,3,4]) = %d, nums = %v\n", k, nums1[:k])

	nums2 := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums2)
	fmt.Printf("MoveZeroes([0,1,0,3,12])                = %v\n", nums2)

	nums3 := []int{3, 2, 2, 3}
	k3 := RemoveElement(nums3, 3)
	fmt.Printf("RemoveElement([3,2,2,3], 3)             = %d, nums = %v\n", k3, nums3[:k3])

	fmt.Println()
	fmt.Println("--- Three Pointers ---")
	fmt.Printf("ThreeSum([-1,0,1,2,-1,-4]) = %v\n", ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("ThreeSum([0,0,0])          = %v\n", ThreeSum([]int{0, 0, 0}))
	fmt.Printf("ThreeSum([1,2,3])          = %v\n", ThreeSum([]int{1, 2, 3}))

	colors := []int{2, 0, 2, 1, 1, 0}
	SortColors(colors)
	fmt.Printf("SortColors([2,0,2,1,1,0])  = %v\n", colors)
}
