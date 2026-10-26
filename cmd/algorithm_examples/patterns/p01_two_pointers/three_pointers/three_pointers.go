package main

import "sort"

// Паттерн: три указателя (three pointers).
//
// Используется, когда нужно выделить три «позиции» в массиве:
// обычно один фиксированный указатель + два сближающихся на
// оставшейся части. Либо три указателя, делящих массив на
// три области (Dutch National Flag).
//
// Сложность: обычно O(n²) или O(n) с сортировкой.

// ThreeSum возвращает все уникальные тройки элементов,
// сумма которых равна нулю. Тройки отсортированы по возрастанию,
// результат не содержит дубликатов.
//
// Сложность: O(n²) время, O(1) доп. память (не считая результата).
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)
	if n < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			switch {
			case sum < 0:
				left++
			case sum > 0:
				right--
			default:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return result
}

// SortColors сортирует массив из 0, 1, 2 in-place за один проход.
// Классический Dutch National Flag problem (Эдсгер Дейкстра).
//
// Идея: три указателя:
//
//	low  — граница зоны 0-ек;
//	mid  — текущий элемент;
//	high — граница зоны 2-ек.
//
// Сложность: O(n) время, O(1) память.
func SortColors(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
