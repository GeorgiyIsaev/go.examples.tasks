package main

import "fmt"

func main() {
	fmt.Println("=== 05. Hash Map / Counters (счётчики) ===")
	fmt.Println()

	freq := CharFrequency("banana")
	fmt.Println("CharFrequency(\"banana\") = ", freq)
	fmt.Println()

	fmt.Println("MajorityElement([2,2,1,1,1,2,2]) =", MajorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // 2
	fmt.Println()

	fmt.Println("TopKFrequent([1,1,1,2,2,3], 2) =", TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)) // [1 2]
	fmt.Println("TopKFrequent([1], 1)           =", TopKFrequent([]int{1}, 1))                // [1]
}
