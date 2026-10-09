package main

import "fmt"

//Спиральный обход матрицы
//Условие: Дана матрица m × n ([][]int).
//Вернуть все её элементы в порядке обхода по спирали
//по часовой стрелке, начиная с верхнего левого угла.
//Примеры:
//•	[[1,2,3],[4,5,6],[7,8,9]] → [1,2,3,6,9,8,7,4,5]
//•	[[1,2,3,4],[5,6,7,8],[9,10,11,12]] → [1,2,3,4,8,12,11,10,9,5,6,7]
//•	[[1]] → [1]
//•	[] → []
//Сигнатура:
//func SpiralOrder(matrix [][]int) []int

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		//1 Верхняя строка слева на право
		for j := left; j <= right; j++ {
			result = append(result, matrix[top][j])
		}
		top++

		//2Правый столбец
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		//3Нижняя строка
		if top <= bottom {
			//(только если есть что обходить)
			for j := right; j >= left; j-- {
				result = append(result, matrix[bottom][j])
			}
			bottom--
		}

		//4 Левый столбец
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

func main() {
	fmt.Println(SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// [1 2 3 6 9 8 7 4 5]

	fmt.Println(SpiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
	// [1 2 3 4 8 12 11 10 9 5 6 7]

	fmt.Println(SpiralOrder([][]int{{1}}))
	// [1]

	fmt.Println(SpiralOrder([][]int{{1, 2, 3}}))
	// [1 2 3]

	fmt.Println(SpiralOrder([][]int{{1}, {2}, {3}}))
	// [1 2 3]

	fmt.Println(SpiralOrder([][]int{}))
	// []
}
