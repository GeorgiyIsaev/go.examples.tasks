package main

import "sort"

// Паттерн: отрезки (intervals) — сортировка + линейный проход.
//
// Базовый приём для ВСЕХ задач на отрезки: сортируем по НАЧАЛУ
// и идём слева направо, поддерживая «текущий объединённый отрезок».
// Следующий отрезок либо расширяет текущий (`start <= curEnd`),
// либо начинает новый (`start > curEnd`).
//
// Сложность: O(n log n) на сортировку, O(1)/O(n) память сверх ответа.
// Важно: НЕ модифицируем входную срезку «на месте» — копируем перед
// сортировкой (иначе ломаются тесты и ожидания вызывающего).

// Merge — объединить все пересекающиеся отрезки (LeetCode 56).
// Касание на границе (`a.end == b.start`) считается пересечением —
// замыкание [1,4] и [4,5] даёт [1,5].
//
// Пример: [[1,3],[2,6],[8,10],[15,18]] -> [[1,6],[8,10],[15,18]].
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sorted := make([][]int, len(intervals))
	for i, iv := range intervals {
		sorted[i] = []int{iv[0], iv[1]}
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][0] < sorted[j][0] })

	result := [][]int{{sorted[0][0], sorted[0][1]}}
	for _, iv := range sorted[1:] {
		last := result[len(result)-1]
		if iv[0] <= last[1] { // пересекается/касается — расширяем
			if iv[1] > last[1] {
				last[1] = iv[1]
			}
		} else { // отдельный отрезок
			result = append(result, []int{iv[0], iv[1]})
		}
	}
	return result
}

// Insert — вставить новый отрезок в отсортированный НЕпересекающийся
// список, объединив пересечения (LeetCode 57).
//
// Три зоны (вход уже отсортирован и не пересекается):
//  1. новый отрезок целиком ЛЕВЕЕ текущего (newEnd < cur.start) —
//     все оставшиеся просто дописываем;
//  2. целиком ПРАВЕЕ (newStart > cur.end) — текущий бережём как есть;
//  3. пересекается — растягиваем newInterval (min/max), текущий
//     съедается.
//
// Пример: [[1,3],[6,9]], new [2,5] -> [[1,5],[6,9]].
func Insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	start, end := newInterval[0], newInterval[1]

	// Зоны 1 и 2: всё, что слева.
	i := 0
	for i < len(intervals) && intervals[i][1] < start {
		result = append(result, intervals[i])
		i++
	}
	// Зона 3: слияние всех пересекающихся.
	for i < len(intervals) && intervals[i][0] <= end {
		if intervals[i][0] < start {
			start = intervals[i][0]
		}
		if intervals[i][1] > end {
			end = intervals[i][1]
		}
		i++
	}
	result = append(result, []int{start, end})
	// Хвост: всё, что справа.
	result = append(result, intervals[i:]...)
	return result
}
