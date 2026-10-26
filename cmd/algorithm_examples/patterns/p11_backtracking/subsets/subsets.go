package main

import "sort"

// Паттерн: backtracking на выбор элемента ("включить / не включить"
// и "взять i-й или нет").
//
// Общий скелет:
//
//	рекурсия(start):
//	    зафиксировать текущее состояние как ответ (если оно валидно)
//	    для i := start; i < n; i++:
//	        выбрать nums[i]          // push
//	        рекурсия(i + 1)          // идём глубже
//	        отменить выбор nums[i]   // pop (backtrack!)
//
// Ключевая деталь — **откат**: строка после рекурсивного вызова
// обязана вернуть состояние ровно таким, каким было. Иначе ветвления
// «протекают» друг в друга и получаются мусорные комбинации.
//
// Subsets — все подмножества (2^n штук), Subsets II — с дублями
// без повторов (сортировка + пропуск одинаковых на одном уровне).

// Subsets — все подмножества nums (LeetCode 78).
// Порядок ответа не важен.
//
// Пример: [1,2,3] -> [[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]].
func Subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		// Каждый вызов — валидное подмножество: копируем path,
		// т.к. он будет меняться в соседних ветках.
		copied := make([]int, len(path))
		copy(copied, path)
		result = append(result, copied)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // выбрать
			backtrack(i + 1)             // углубиться
			path = path[:len(path)-1]    // откатить
		}
	}
	backtrack(0)
	return result
}

// SubsetsWithDup — все подмножества без повторов при дублях в nums
// (LeetCode 90).
//
// Трюк: сортируем, чтобы одинаковые стояли рядом, и на КАЖДОМ уровне
// рекурсии пропускаем i, если nums[i] == nums[i-1] и i > start —
// иначе одно и то же подмножество соберётся разными путями.
//
// Пример: [1,2,2] -> [[], [1], [1,2], [1,2,2], [2], [2,2]].
func SubsetsWithDup(nums []int) [][]int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	result := [][]int{}
	path := []int{}

	var backtrack func(start int)
	backtrack = func(start int) {
		copied := make([]int, len(path))
		copy(copied, path)
		result = append(result, copied)

		for i := start; i < len(sorted); i++ {
			// Пропуск дублей на одном и том же уровне.
			if i > start && sorted[i] == sorted[i-1] {
				continue
			}
			path = append(path, sorted[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return result
}
