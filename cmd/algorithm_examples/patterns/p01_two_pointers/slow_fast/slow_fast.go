package main

// Паттерн: два указателя в одном направлении (slow / fast).
//
// Идея: fast бежит впереди и читает, slow плетётся сзади и пишет.
// Часто используется для in-place модификации среза, когда
// нужно «сжать» данные, выкинув часть элементов.
//
// Также это основа для алгоритма Floyd (черепаха и заяц) для
// поиска цикла, но там структура — связный список.
//
// Сложность: O(n) время, O(1) память.

// RemoveDuplicatesSorted удаляет дубликаты из отсортированного
// среза in-place. Возвращает новую длину. Элементы после новой
// длины значения не имеют.
//
// Пример: [1,1,2,3,3,3,4] -> длина 4, первые 4 элемента = [1,2,3,4].
func RemoveDuplicatesSorted(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// MoveZeroes перемещает все нули в конец среза in-place,
// сохраняя порядок ненулевых элементов.
//
// Пример: [0,1,0,3,12] -> [1,3,12,0,0].
func MoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

// RemoveElement удаляет все вхождения val из среза in-place
// и возвращает новую длину.
//
// Пример: [3,2,2,3], val=3 -> длина 2, nums[:2] = [2,2].
func RemoveElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
