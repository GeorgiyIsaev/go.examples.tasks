package main

import "fmt"

func main() {
	fmt.Println("=== 22. Bit Manipulation / Basics ===")
	fmt.Println()

	fmt.Println("HammingWeight(11) =", HammingWeight(11)) // 3 (1011)
	fmt.Println("HammingWeight(255) =", HammingWeight(255))
	fmt.Println()

	for _, n := range []int{1, 2, 16, 18, 0, -2} {
		fmt.Printf("IsPowerOfTwo(%d) = %v\n", n, IsPowerOfTwo(n))
	}
	fmt.Println()

	fmt.Println("CountBits(5) =", CountBits(5)) // [0 1 1 2 1 2]
}
