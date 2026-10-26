package main

import "fmt"

// Запуск: go run ./middle_find
func main() {
	fmt.Println("=== 02. Fast / Slow / Middle Find (середина списка) ===")
	fmt.Println()

	odd := BuildList([]int{1, 2, 3, 4, 5})
	even := BuildList([]int{1, 2, 3, 4})
	fmt.Printf("MiddleNode([1,2,3,4,5]).Val          = %d (ожидается 3)\n", MiddleNode(odd).Val)
	fmt.Printf("MiddleNode([1,2,3,4]).Val            = %d (ожидается 3, второй средний)\n",
		MiddleNode(even).Val)

	fmt.Println()

	fmt.Printf("IsPalindromeList([1,2,2,1])          = %v\n",
		IsPalindromeList(BuildList([]int{1, 2, 2, 1})))
	fmt.Printf("IsPalindromeList([1,2,3,2,1])        = %v\n",
		IsPalindromeList(BuildList([]int{1, 2, 3, 2, 1})))
	fmt.Printf("IsPalindromeList([1,2])              = %v\n",
		IsPalindromeList(BuildList([]int{1, 2})))
	fmt.Printf("IsPalindromeList([1])                = %v\n",
		IsPalindromeList(BuildList([]int{1})))
}
