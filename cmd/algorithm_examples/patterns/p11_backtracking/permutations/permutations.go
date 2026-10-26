package main

import "sort"

// Backtracking: перестановки — выбор на КАЖДОЙ позиции результата.
//
// Отличие от subsets: там мы шли по входу слева направо и решали
// «взять i-й элемент или нет» (start защищает от переиспользования).
// Здесь результат — переполненная перестановка: на позицию k берём
// любой ещё не использованный элемент, а `used[]` не даёт взять его
// снова.
//
//	рекурсия():
//	    len(path) == n -> ответ
//	    для i := 0; i < n; i++:
//	        если used[i] -> пропуск
//	        path = append(path, nums[i]); used[i] = true
//	        рекурсия()
//	        path = path[:len-1]; used[i] = false   // откат

// Permute — все перестановки уникальных элементов (LeetCode 46).
//
// Пример: [1,2,3] -> 6 перестановок.
func Permute(nums []int) [][]int {
	result := [][]int{}
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			copied := make([]int, len(nums))
			copy(copied, path)
			result = append(result, copied)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1] // откат
			used[i] = false
		}
	}
	backtrack()
	return result
}

// PermuteUnique — перестановки без повторов при дублях в nums
// (LeetCode 47).
//
// Трюк: сортируем и на одном уровне рекурсии пропускаем i, если
// nums[i] == nums[i-1] и nums[i-1] ещё НЕ использована в этой ветке
// (тогда дубль уже «отработал» бы тем же путём раньше).
//
// Пример: [1,1,2] -> [[1,1,2],[1,2,1],[2,1,1]].
func PermuteUnique(nums []int) [][]int {
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	sort.Ints(sorted)

	result := [][]int{}
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			copied := make([]int, len(nums))
			copy(copied, path)
			result = append(result, copied)
			return
		}
		for i := 0; i < len(sorted); i++ {
			if used[i] {
				continue
			}
			// Дубль пропускаем, только если предыдущий такой же
			// элемент ещё висит на этом уровне (не откатнут).
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, sorted[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return result
}
