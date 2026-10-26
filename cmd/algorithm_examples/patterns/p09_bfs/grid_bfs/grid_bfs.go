package main

// Паттерн: BFS по сетке (grid BFS).
//
// Идея: та же очередь, что и в графе, только «узлы» — клетки сетки,
// а «соседи» — 4 (или 8) окружающие клетки. Каждую посещённую клетку
// помечаем сразу, при постановке в очередь, а не при извлечении —
// иначе одна и та же клетка попадёт в очередь несколько раз.
//
// Когда применять:
//   - «острова», «заливка», «компоненты в сетке»;
//   - «минимальное число шагов по сетке»;
//   - «много источников одновременно» (multi-source BFS).
//
// Сложность: O(rows·cols) время и память.

// NumIslands — количество островов (4-связность, 1 — суша, 0 — вода).
// После обхода сетка не меняется (работаем с копией посещённости).
//
// Пример:
//
//	1 1 0     -> 2 острова
//	1 0 0
//	0 0 1
func NumIslands(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 || visited[r][c] {
				continue
			}
			// Нашли новый остров — заливаем его целиком.
			count++
			visited[r][c] = true
			queue := [][2]int{{r, c}}
			for len(queue) > 0 {
				cr, cc := queue[0][0], queue[0][1]
				queue = queue[1:]
				for _, d := range dirs {
					nr, nc := cr+d[0], cc+d[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
						continue
					}
					if grid[nr][nc] != 1 || visited[nr][nc] {
						continue
					}
					visited[nr][nc] = true // помечаем при постановке
					queue = append(queue, [][2]int{{nr, nc}}...)
				}
			}
		}
	}
	return count
}

// RottingOranges — минут, за которые все свежие апельсины сгниют.
// За минуту гнилой апельсин заражает 4 соседей. −1, если сгнить
// не всем.
//
// Классический multi-source BFS: в очередь кладём ВСЕ гнилые
// апельсины сразу, каждая итерация = одна минута.
//
// Пример: [[2,1,1],[1,1,0],[0,1,1]] -> 4.
func RottingOranges(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	queue := [][2]int{}
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case 2:
				queue = append(queue, [][2]int{{r, c}}...)
			case 1:
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}

	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		// Все узлы текущей минуты.
		for size := len(queue); size > 0; size-- {
			cr, cc := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := cr+d[0], cc+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				fresh--
				queue = append(queue, [][2]int{{nr, nc}}...)
			}
		}
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

// ShortestPathBinaryMatrix — длина кратчайшего пути из (0,0) в (n-1,n-1)
// в бинарной сетке (0 — проходимо, 1 — стена), движение по 8 направлениям.
// −1, если пути нет.
//
// Пример: [[0,0,0],[1,1,0],[1,1,0]] -> 4.
func ShortestPathBinaryMatrix(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	dirs := [][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	type cell struct{ r, c int }
	queue := []cell{{0, 0}}
	visited[0][0] = true
	steps := 1

	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			cur := queue[0]
			cr, cc := cur.r, cur.c
			queue = queue[1:]
			for _, d := range dirs {
				nr, nc := cr+d[0], cc+d[1]
				if nr == n-1 && nc == n-1 {
					return steps + 1
				}
				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] == 1 || visited[nr][nc] {
					continue
				}
				visited[nr][nc] = true
				queue = append(queue, cell{nr, nc})
			}
		}
		steps++
	}
	return -1
}
