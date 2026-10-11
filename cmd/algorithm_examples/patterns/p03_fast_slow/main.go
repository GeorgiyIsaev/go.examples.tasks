package main

import "fmt"

func main() {
	fmt.Println("=== 03. Fast / Slow (Floyd's algorithm) ===")
	fmt.Println()

	fmt.Println("--- Cycle Detection ---")

	// Список без цикла: 1 -> 2 -> 3 -> 4
	noCycle := BuildList([]int{1, 2, 3, 4})
	fmt.Printf("HasCycle(1->2->3->4)                = %v\n", HasCycle(noCycle))
	fmt.Printf("DetectCycle(1->2->3->4)             = %v\n", DetectCycle(noCycle))
	fmt.Printf("CountCycleLength(1->2->3->4)        = %d\n", CountCycleLength(noCycle))

	// Список с циклом: 3 -> 2 -> 0 -> -4 -> (обратно на 2)
	cyc := BuildList([]int{3, 2, 0, -4})
	tail := findNodeByIndex(cyc, 3)
	entry := findNodeByIndex(cyc, 1) // узел со значением 2
	tail.Next = entry

	fmt.Printf("HasCycle(3->2->0->-4->2...)         = %v\n", HasCycle(cyc))
	fmt.Printf("DetectCycle(...).Val                = %d (ожидается 2)\n", DetectCycle(cyc).Val)
	fmt.Printf("CountCycleLength(...)               = %d (ожидается 3)\n", CountCycleLength(cyc))

	fmt.Println()
	fmt.Println("--- Find Duplicate ---")
	fmt.Printf("FindDuplicate([1,3,4,2,2])          = %d\n", FindDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Printf("FindDuplicate([3,1,3,4,2])          = %d\n", FindDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Printf("FindDuplicate([1,1])                = %d\n", FindDuplicate([]int{1, 1}))

	fmt.Println()
	fmt.Println("--- Middle & Palindrome ---")

	odd := BuildList([]int{1, 2, 3, 4, 5})
	even := BuildList([]int{1, 2, 3, 4})
	fmt.Printf("MiddleNode([1,2,3,4,5]).Val         = %d (ожидается 3)\n", MiddleNode(odd).Val)
	fmt.Printf("MiddleNode([1,2,3,4]).Val           = %d (ожидается 3, второй средний)\n", MiddleNode(even).Val)

	fmt.Printf("IsPalindromeList([1,2,2,1])         = %v\n",
		IsPalindromeList(BuildList([]int{1, 2, 2, 1})))
	fmt.Printf("IsPalindromeList([1,2])             = %v\n",
		IsPalindromeList(BuildList([]int{1, 2})))
	fmt.Printf("IsPalindromeList([1])               = %v\n",
		IsPalindromeList(BuildList([]int{1})))

	fmt.Println()
	fmt.Println("--- Happy Number ---")
	fmt.Printf("IsHappy(19)                          = %v\n", IsHappy(19))
	fmt.Printf("IsHappy(2)                           = %v\n", IsHappy(2))
	fmt.Printf("IsHappy(1)                           = %v\n", IsHappy(1))
}
