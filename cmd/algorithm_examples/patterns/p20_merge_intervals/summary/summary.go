package main

// Summary Ranges (LeetCode 228): отсортированный массив УНИКАЛЬНЫХ
// чисел → список строковых диапазонов подряд идущих.
//
// Линейный проход «где начинается и где кончивается серия»:
//   - идём, пока `nums[i] == start + (i - startIdx)`, т.е. серия
//     непрерывна (в отсортированном массиве — просто `nums[i+1] == nums[i]+1`);
//   - серия закончилась — записываем "start->end" или "start",
//     если длина 1.
//
// Сложность: O(n) — сортировка уже есть по условию.
//
// Пример: [0,1,2,4,5,7] -> ["0->2","4->5","7"].

// SummaryRanges возвращает строковые диапазоны.
func SummaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}

	start := nums[0] // начало текущей серии
	for i := 1; i <= len(nums); i++ {
		// Серия кончилась: конец массива или разрыв (+1).
		if i == len(nums) || nums[i] != nums[i-1]+1 {
			end := nums[i-1]
			if start == end {
				result = append(result, itoa(start))
			} else {
				result = append(result, itoa(start)+"->"+itoa(end))
			}
			if i < len(nums) {
				start = nums[i]
			}
		}
	}
	return result
}

// itoa — минимальный конвертер int -> строка (без strconv в demo-коде).
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	negative := n < 0
	if negative {
		n = -n
	}
	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	if negative {
		return "-" + string(digits)
	}
	return string(digits)
}
