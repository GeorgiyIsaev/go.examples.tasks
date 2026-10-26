package main

import "fmt"

func main() {
	fmt.Println("=== 11. Backtracking / Permutations (все перестановки) ===")
	fmt.Println()

	perms := Permute([]int{1, 2, 3})
	fmt.Println("Permute([1,2,3]):")
	for _, p := range perms {
		fmt.Println(" ", p)
	}
	fmt.Println("всего:", len(perms), "(3!)")
	fmt.Println()

	fmt.Println("PermuteUnique([1,1,2]):")
	for _, p := range PermuteUnique([]int{1, 1, 2}) {
		fmt.Println(" ", p)
	}
	fmt.Println()

	fmt.Println("Permute([]) =", Permute([]int{}))
}
