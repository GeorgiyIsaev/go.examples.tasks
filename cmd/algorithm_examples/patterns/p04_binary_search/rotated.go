package main

// Паттерн: бинарный поиск в повёрнутом (rotated) отсортированном массиве.
//
// Массив получен из отсортированного сдвигом влево на k позиций.
// Пример: [0,1,2,4,5,6,7] сдвинули -> [4,5,6,7,0,1,2].
//
// Ключевая идея: даже в повёрнутом массиве хотя бы одна половина
// (относительно mid) всегда ОТСОРТИРОВАНА. По ней мы можем судить,
// в какой половине искать target.

// SearchRotated — индекс target в повёрнутом массиве или -1.
//
// Условие: все элементы уникальны.
//
// Идея:
//  1. Смотрим на nums[mid].
//  2. Если nums[left] <= nums[mid] — левая половина отсортирована.
//     Тогда: если target в [nums[left], nums[mid]) — идём влево,
//     иначе вправо.
//  3. Иначе правая половина отсортирована. Аналогично.
func SearchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// Левая половина отсортирована.
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Правая половина отсортирована.
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// FindMinRotated — минимальный элемент в повёрнутом массиве.
//
// Идея: сравниваем nums[mid] с nums[right]. Если nums[mid] > nums[right],
// минимум СПРАВА от mid. Иначе — слева (включая mid).
// Тот же приём работает с дубликатами, если немного модифицировать
// (при nums[mid] == nums[right] делать right--).
func FindMinRotated(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
