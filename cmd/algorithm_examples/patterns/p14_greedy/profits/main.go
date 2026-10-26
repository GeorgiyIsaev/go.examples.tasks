package main

import "fmt"

func main() {
	fmt.Println("=== 14. Greedy / Profits (купил-продал) ===")
	fmt.Println()

	for _, p := range [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{1, 2},
		{},
	} {
		fmt.Printf("MaxProfit(%v) = %d | MaxProfitMany = %d\n",
			p, MaxProfit(p), MaxProfitMany(p))
	}
}
