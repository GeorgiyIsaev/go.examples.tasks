package main

// Backtracking на доске: «заполнить клетку — проверить конфликт —
// откатить».
//
// Отличие от permutations/subsets: состояние — не срез, а доска,
// а кандидатов на клетку несколько (ферзи, буквы). Главная сложность —
// дешёвая проверка конфликта: вместо полного осмотра доски держим
// множества «заняты ли диагональ/столбец/ряд» и обновляем их за O(1)
// при выборе и откате.
//
//	рекурсия(row):
//	    row == n -> записать решение
//	    для col := 0..n-1:
//	        если конфликт -> пропуск
//	        поставить ферзя (обновить sets)
//	        рекурсия(row + 1)
//	        убрать ферзя (откатить sets)

// SolveNQueens — все расстановки N ферзей на доске N×N (LeetCode 51).
// Возвращает каждый вариант строкой вида ".Q..":
//
//	[".Q..","...Q","Q...","...Q"]
//
// Сложность: O(n!) в худшем — на каждом уровне выбор сужается.
func SolveNQueens(n int) [][]string {
	result := [][]string{}
	cols := map[int]bool{}  // занятые столбцы
	diag1 := map[int]bool{} // диагонали r-c (константа на диагонали)
	diag2 := map[int]bool{} // диагонали r+c
	board := make([][]byte, n)
	for i := range board {
		board[i] = bytesRepeat('.', n)
	}

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			solution := make([]string, n)
			for i, line := range board {
				solution[i] = string(line)
			}
			result = append(result, solution)
			return
		}
		for col := 0; col < n; col++ {
			// Конфликт по столбцу или диагонали?
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}
			// Поставить ферзя.
			cols[col], diag1[row-col], diag2[row+col] = true, true, true
			board[row][col] = 'Q'

			backtrack(row + 1)

			// Откат.
			board[row][col] = '.'
			cols[col], diag1[row-col], diag2[row+col] = false, false, false
		}
	}
	backtrack(0)
	return result
}

// CountNQueens — сколько решений у N-Queens (LeetCode 52).
// n=4 -> 2, n=8 -> 92.
func CountNQueens(n int) int {
	return len(SolveNQueens(n))
}

// ExistWord — есть ли слово в сетке буква к букве, соседними по
// 4 сторонам, с возвратом по тому же пути (LeetCode 79).
//
// Ключевой трюк: клетку, в которую зашли, помечаем (например, '#'),
// чтобы обойтись без отдельного visited и не зациклиться; после
// рекурсии возвращаем букву — это и есть откат.
//
// Пример: board=[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],
// word="ABCCED" -> true.
func ExistWord(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])

	var dfs func(r, c, k int) bool
	dfs = func(r, c, k int) bool {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}
		if board[r][c] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true // всё слово совпало
		}
		saved := board[r][c]
		board[r][c] = '#' // помечаем посещённой

		found := dfs(r+1, c, k+1) ||
			dfs(r-1, c, k+1) ||
			dfs(r, c+1, k+1) ||
			dfs(r, c-1, k+1)

		board[r][c] = saved // откат
		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

// bytesRepeat — заполнитель доски, чтобы не тянуть strings.
func bytesRepeat(b byte, n int) []byte {
	out := make([]byte, n)
	for i := range out {
		out[i] = b
	}
	return out
}
