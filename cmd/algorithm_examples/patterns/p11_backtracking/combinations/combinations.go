package main

import "sort"

// Backtracking: комбинации — «выбрать k из n» и «сумма в цель».
//
// Отличие от subsets тем, что нас ограничивает ЦЕЛЬ (k штук или
// сумма target) — рекурсия заканчивается не по всему входу, а по
// условию остановки, и мы можем обрезать ветку досрочно
// (pruning): если сумма уже переросла цель — вглубь не идём.
//
//	рекурсия(start, current):
//	    условие остановки -> записать ответ
//	    для i := start; i < n; i++ (с пропусками/обрезками):
//	        выбрать nums[i]
//	        рекурсия(i + 1 или i)   // i+1 — без повторов, i — с повтором
//	        откатить

// Combine — все комбинации из n чисел длины k (LeetCode 77).
// Числа 1..n.
//
// Пример: n=4, k=2 -> [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]].
func Combine(n, k int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path) == k {
			copied := make([]int, k)
			copy(copied, path)
			result = append(result, copied)
			return
		}
		// Обрезка: сколько элементов ещё можем взять?
		// Если свободных мест (n-i+1) меньше, чем не хватает до k —
		// эта ветка не соберёт комбинацию нужной длины.
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return result
}

// CombinationSum — комбинации из candidates с повторным использованием
// элемента, дающие сумму target (LeetCode 39). Каждый элемент можно
// брать сколько угодно раз.
//
// Правило индекса: при разрешённом повторе рекурсия идёт с ТЕМ ЖЕ i —
// тогда nums[i] можно взять ещё раз. Стартовый цикл — с i, а не i+1.
//
// Пример: [2,3,6,7], 7 -> [[2,2,3],[7]].
func CombinationSum(candidates []int, target int) [][]int {
	// Сортируем, чтобы можно было обрезать ветку: как только
	// элемент больше остатка — все следующие ещё больше.
	sorted := make([]int, len(candidates))
	copy(sorted, candidates)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}

	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			copied := make([]int, len(path))
			copy(copied, path)
			result = append(result, copied)
			return
		}
		for i := start; i < len(sorted); i++ {
			if sorted[i] > remain {
				break // дальше только больше — ветка мертва
			}
			path = append(path, sorted[i])
			backtrack(i, remain-sorted[i]) // i, не i+1: можно повторить
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return result
}

// CombinationSumII — как CombinationSum, но каждый элемент можно взять
// ровно один раз, в candidates есть дубли (LeetCode 40).
//
// Сортируем + пропускаем дубль на одном уровне (как в subsets II).
// Рекурсия — с i+1: повторов не разрешаем.
//
// Пример: [10,1,2,7,6,1,5], 8 -> [[1,1,6],[1,2,5],[1,7],[2,6]].
func CombinationSumII(candidates []int, target int) [][]int {
	sorted := make([]int, len(candidates))
	copy(sorted, candidates)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}

	var backtrack func(start, remain int)
	backtrack = func(start, remain int) {
		if remain == 0 {
			copied := make([]int, len(path))
			copy(copied, path)
			result = append(result, copied)
			return
		}
		for i := start; i < len(sorted); i++ {
			if sorted[i] > remain {
				break // дальше только больше — выходим
			}
			if i > start && sorted[i] == sorted[i-1] {
				continue // дубль на этом уровне
			}
			path = append(path, sorted[i])
			backtrack(i+1, remain-sorted[i])
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target)
	return result
}
