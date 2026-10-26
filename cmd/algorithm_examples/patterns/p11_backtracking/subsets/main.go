package main

import "fmt"

func main() {
	fmt.Println("=== 11. Backtracking / Subsets (все подмножества) ===")
	fmt.Println()

	fmt.Println("Subsets([1,2,3]):")
	for _, s := range Subsets([]int{1, 2, 3}) {
		fmt.Println(" ", s)
	}
	fmt.Println("всего подмножеств:", len(Subsets([]int{1, 2, 3})), "(2^3)")
	fmt.Println()

	fmt.Println("SubsetsWithDup([1,2,2]):")
	for _, s := range SubsetsWithDup([]int{1, 2, 2}) {
		fmt.Println(" ", s)
	}
	fmt.Println()

	fmt.Println("Subsets([]) =", Subsets([]int{}))
}
