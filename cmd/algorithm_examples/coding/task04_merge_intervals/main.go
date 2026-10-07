package main

import (
	"fmt"
	"sort"
)

//Слияние интервалов
//Условие: Дан [][]int интервалов вида [start, end].
//Слить пересекающиеся интервалы, вернуть отсортированный по началу результат.
//Пример:
//[[1,3],[2,6],[8,10],[15,18]] -> [[1,6],[8,10],[15,18]]
//Сигнатура:
//func MergeIntervals(intervals [][]int) [][]int

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 1. Сортируем по началу интервала
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals) //копируем
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	// 2. Сливаем
	result := [][]int{sorted[0]} // кладём первый интервал

	for i := 1; i < len(sorted); i++ {
		last := result[len(result)-1] // последний в результате
		cur := sorted[i]

		if last[1] < cur[0] {
			result = append(result, cur)
		} else {
			result[len(result)-1][1] = cur[1]
		}
	}

	return result
}

func main() {

	sl := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(sl)
	fmt.Println(MergeIntervals(sl))

	sl = [][]int{{1, 3}, {12, 16}, {2, 10}, {15, 18}}
	fmt.Println(sl)
	fmt.Println(MergeIntervals(sl))

	sl = [][]int{{1, 3}, {12, 16}, {2, 12}, {15, 18}}
	fmt.Println(sl)
	fmt.Println(MergeIntervals(sl))

}
