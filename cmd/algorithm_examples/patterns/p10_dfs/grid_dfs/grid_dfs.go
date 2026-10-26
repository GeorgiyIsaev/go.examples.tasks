package main

// Паттерн: DFS по сетке (заливка / flood fill).
//
// Идея: вместо очереди — рекурсия (или явный стек). Нашли стартовую
// клетку — «заливаем» всю область, помечая посещённые клетки сразу,
// чтобы не зайти повторно и не зациклиться.
//
// ВАЖНО: у grid n·m клеток и глубина рекурсии может дойти до n·m —
// на больших сетках рекурсия упрётся в лимит стека, тогда используй
// явный стек или BFS (папка p09_bfs/grid_bfs).
//
// Сложность: O(rows·cols) время и память.

// длины 4 направлений (вверх, вниз, влево, вправо).
var gridDirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

// NumIslandsDFS — количество островов через DFS-заливку
// (1 — суша, 0 — вода, 4-связность).
//
// Аналог NumIslands из p09_bfs, но вместо очереди — рекурсия.
//
// Пример:
//
//	1 1 0     -> 2 острова
//	1 0 0
//	0 0 1
func NumIslandsDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var flood func(r, c int)
	flood = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if grid[r][c] != 1 || visited[r][c] {
			return
		}
		visited[r][c] = true
		for _, d := range gridDirs {
			flood(r+d[0], c+d[1])
		}
	}

	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				count++
				flood(r, c)
			}
		}
	}
	return count
}

// MaxAreaOfIsland — площадь самого большого острова
// (LeetCode 695, 1 — суша).
//
// Пример: [[0,0,1],[1,1,1],[1,0,0],[1,1,1]] -> 5.
func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	var area func(r, c int) int
	area = func(r, c int) int {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return 0
		}
		if grid[r][c] != 1 || visited[r][c] {
			return 0
		}
		visited[r][c] = true
		total := 1
		for _, d := range gridDirs {
			total += area(r+d[0], c+d[1])
		}
		return total
	}

	maxArea := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 && !visited[r][c] {
				if a := area(r, c); a > maxArea {
					maxArea = a
				}
			}
		}
	}
	return maxArea
}

// FloodFill — заливка (LeetCode 733): меняем цвет старой области
// на новый цвет, начиная из (sr, sc). Возвращает новую сетку.
//
// Пример: image=[[1,1,1],[1,1,0],[1,0,1]], (1,1), 2 ->
//
//	[[2,2,2],[2,2,0],[2,0,1]]
func FloodFill(image [][]int, sr, sc, newColor int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image // иначе зациклится: все клетки уже «новые»
	}
	rows, cols := len(image), len(image[0])

	var fill func(r, c int)
	fill = func(r, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		if image[r][c] != oldColor {
			return
		}
		image[r][c] = newColor
		for _, d := range gridDirs {
			fill(r+d[0], c+d[1])
		}
	}
	fill(sr, sc)
	return image
}
