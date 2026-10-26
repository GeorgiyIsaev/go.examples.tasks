package main

import "fmt"

func main() {
	fmt.Println("=== 11. Backtracking / Board (доска: ферзи и слова) ===")
	fmt.Println()

	solutions := SolveNQueens(4)
	fmt.Println("N-Queens n=4, решений:", len(solutions))
	for _, s := range solutions {
		fmt.Println(" ---")
		for _, line := range s {
			fmt.Println(" ", line)
		}
	}
	fmt.Println()
	fmt.Println("CountNQueens(8) =", CountNQueens(8), "(должно быть 92)")
	fmt.Println()

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println("доска A B C E / S F C S / A D E E")
	fmt.Println("word \"ABCCED\"  ->", ExistWord(board, "ABCCED")) // true
	fmt.Println("word \"SEE\"    ->", ExistWord(board, "SEE"))     // true
	fmt.Println("word \"ABCB\"   ->", ExistWord(board, "ABCB"))    // false
}
