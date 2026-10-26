package main

// Паттерн: границы (lower bound, upper bound).
//
// Классический бинарный поиск возвращает индекс ТОЧНОГО совпадения.
// Но часто нужно не совпадение, а ГРАНИЦА:
//   - первый элемент >= target (lower bound);
//   - первый элемент > target (upper bound).
//
// Через эти два приёма решается большинство «хитрых» задач на поиск.
// Шаблон: right = n (не n-1!), условие left < right, mid всегда
// двигает одну из границ к mid (не к mid ± 1 с обеих сторон).

// LowerBound — первый индекс i, где nums[i] >= target.
// Если все элементы < target — вернёт len(nums).
//
// Пример: [1,3,3,5,7], target=3 -> 1.
//
//	[1,3,3,5,7], target=4 -> 3.
//	[1,3,3,5,7], target=8 -> 5.
func LowerBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// UpperBound — первый индекс i, где nums[i] > target.
// Если все элементы <= target — вернёт len(nums).
//
// Пример: [1,3,3,5,7], target=3 -> 3.
//
//	[1,3,3,5,7], target=7 -> 5.
func UpperBound(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// SearchRange — первое и последнее вхождение target.
// Если target нет — вернёт [-1, -1].
//
// Пример: [5,7,7,8,8,10], target=8 -> [3,4].
//
//	[5,7,7,8,8,10], target=6 -> [-1,-1].
func SearchRange(nums []int, target int) []int {
	first := LowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}
	last := UpperBound(nums, target) - 1
	return []int{first, last}
}

// CountOccurrences — сколько раз target встречается в отсортированном
// массиве. O(log n) — не O(n), как при обычном линейном подсчёте.
func CountOccurrences(nums []int, target int) int {
	first := LowerBound(nums, target)
	if first == len(nums) || nums[first] != target {
		return 0
	}
	return UpperBound(nums, target) - first
}
