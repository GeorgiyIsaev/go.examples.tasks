package main

import "fmt"

// Запуск: go run ./min_window
func main() {
	fmt.Println("=== 03. Sliding Window / Min (сжимающееся окно, минимум) ===")
	fmt.Println()

	fmt.Printf("MinSubArrayLen(7, [2,3,1,2,4,3])          = %d\n",
		MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Printf("MinSubArrayLen(100, [1,2,3])              = %d\n",
		MinSubArrayLen(100, []int{1, 2, 3}))
	fmt.Printf("MinSubArrayLen(4, [1,4,4])                = %d\n",
		MinSubArrayLen(4, []int{1, 4, 4}))

	fmt.Printf("MinWindow(\"ADOBECODEBANC\", \"ABC\")            = %q\n",
		MinWindow("ADOBECODEBANC", "ABC"))
	fmt.Printf("MinWindow(\"a\", \"aa\")                        = %q\n",
		MinWindow("a", "aa"))
	fmt.Printf("MinWindow(\"cabwefgewcwaefgcf\", \"cae\")         = %q\n",
		MinWindow("cabwefgewcwaefgcf", "cae"))
}
