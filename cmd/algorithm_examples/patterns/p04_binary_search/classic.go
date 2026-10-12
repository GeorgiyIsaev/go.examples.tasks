package main

// Паттерн: классический бинарный поиск.
//
// Идея: поддерживаем интервал [left, right], в котором точно находится
// ответ. На каждом шаге смотрим в середину и отбрасываем половину,
// где ответа быть не может. O(log n) вместо O(n).
//
// Условие применимости: массив ОТСОРТИРОВАН, и мы умеем за O(1)
// определить, в какой половине искать дальше.

// BinarySearch — индекс target в отсортированном nums или -1.
//
// Инвариант: ответ (если есть) всегда внутри [left, right].
// Условие цикла left <= right: обрабатываем случай одного элемента.
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 // защита от overflow
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

// SearchInsert — индекс, куда нужно вставить target, чтобы массив
// остался отсортированным. Если target уже есть — его индекс.
//
// Пример: [1,3,5,6], target=5 -> 2.
//
//	[1,3,5,6], target=2 -> 1.
//	[1,3,5,6], target=7 -> 4.
//	[1,3,5,6], target=0 -> 0.
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums) // right = n, не n-1
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

// FirstBadVersion — первая «плохая» версия среди 1..n.
// isBadVersion(i) возвращает true для i >= firstBad.
//
// Идея: у нас монотонное пространство — good...good bad...bad.
// Ищем первый bad.
func FirstBadVersion(n int, isBadVersion func(int) bool) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
