package main

import "fmt"

func main() {
	fmt.Println("=== 22. Bit Manipulation / Tricks ===")
	fmt.Println()

	fmt.Println("HammingDistance(1, 4) =", HammingDistance(1, 4)) // 2
	fmt.Println("HammingDistance(3, 3) =", HammingDistance(3, 3)) // 0
	fmt.Println()

	fmt.Printf("ReverseBits(43261596) = %d\n", ReverseBits(43261596))
	fmt.Printf("ReverseBits(1) = %d\n", ReverseBits(1)) // 1<<31 = 2147483648
	fmt.Println()

	for _, p := range [][2]int{{5, 7}, {-1, 1}, {0, 0}, {-3, -4}} {
		fmt.Printf("GetSum(%d, %d) = %d\n", p[0], p[1], GetSum(p[0], p[1]))
	}
}
