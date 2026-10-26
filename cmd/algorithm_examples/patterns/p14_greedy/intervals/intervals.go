package main

import "sort"

// Activity Selection — «выбор мероприятий»: классическая жадность.
//
// Задача: выбрать максимально число НЕПЕРЕСЕКАЮЩИХся отрезков.
// Жадное правило: сортируем по КОНЦУ и всегда берём тот, что
// заканчивается раньше всех оставшихся. Тогда следующий конфликт
// случится как можно позже — останется максимум места для остальных.
//
// Почему это глобально оптимально (аргумент для собеса): пусть
// оптимум берёт первый отрезок R1, а жадный — Rg (конец позже или
// равен). R1 можно заменить на Rg — конец не вырос, значит,
// совместимость с оставшимися не ухудшилась. Повторяя, из любого
// оптимума получаем жадное решение => оно тоже оптимально.
//
// Сложность: O(n log n) (сортировка), O(1) сверх сортировки.

// EraseOverlapIntervals — минимальное число удалений, чтобы отрезки
// стали непересекающимися (LeetCode 435; ответ = n - максимум
// непересекающихся).
//
// Пример: [[1,2],[2,3],[3,4],[1,3]] -> 1 (убрать [1,3]).
func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)
	// Сортировка по КОНЦУ (при равенстве — по началу).
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i][1] != sorted[j][1] {
			return sorted[i][1] < sorted[j][1]
		}
		return sorted[i][0] < sorted[j][0]
	})

	kept := 1
	lastEnd := sorted[0][1]
	for i := 1; i < len(sorted); i++ {
		if sorted[i][0] >= lastEnd { // не пересекается (касание — ок)
			kept++
			lastEnd = sorted[i][1]
		}
	}
	return len(intervals) - kept
}

// FindMinArrowShots — минимальное число «стрел», чтобы пробить все
// шары (LeetCode 452). Шар описан [xstart, xend]; стрела в точке x
// пробивает все шары с xstart <= x <= xend.
//
// Тот же activity selection: сортируем по концу и стреляем в
// КОНЕЦ первого шара — так пробиваем максимум остальных.
//
// Пример: [[10,16],[2,8],[1,6],[7,12]] -> 2.
func FindMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sorted := make([][]int, len(points))
	copy(sorted, points)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i][1] != sorted[j][1] {
			return sorted[i][1] < sorted[j][1]
		}
		return sorted[i][0] < sorted[j][0]
	})

	arrows := 1
	shotAt := sorted[0][1]
	for i := 1; i < len(sorted); i++ {
		if sorted[i][0] > shotAt { // мимо: нужна новая стрела
			arrows++
			shotAt = sorted[i][1]
		}
		// sorted[i][0] <= shotAt => шар пробит этой же стрелой
	}
	return arrows
}
