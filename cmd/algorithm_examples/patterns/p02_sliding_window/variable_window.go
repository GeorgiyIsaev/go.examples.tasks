package main

// Паттерн: растущее окно (variable-size sliding window, поиск максимума).
//
// Идея: окно [left, right] расширяется вправо. Пока условие выполняется —
// просто двигаем right. Как только условие нарушено — двигаем left, пока
// условие снова не станет верным. Ответ обновляем на каждом шаге.
//
// Когда применять:
//   - "самая длинная подстрока/подмассив с условием";
//   - "максимальная длина окна";
//   - "наибольшее количество ... в окне";
//   - свойство МОНОТОННО: при увеличении окна условие может испортиться,
//     при уменьшении — восстановиться.
//
// Сложность: O(n) время (каждый элемент входит и выходит из окна один раз).

// LengthOfLongestSubstring — длина самой длинной подстроки без
// повторяющихся символов. Работает с байтами (ASCII/латиница).
//
// Пример: "abcabcbb" -> 3 ("abc").
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		// Если символ уже встречался внутри окна — двигаем left
		// на позицию сразу после его прошлого вхождения.
		if pos, ok := lastSeen[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastSeen[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// LongestOnes — максимальная длина подмассива из единиц, если можно
// заменить не более k нулей на единицы.
//
// Пример: nums=[1,1,1,0,0,0,1,1,1,1,0], k=2 -> 6.
func LongestOnes(nums []int, k int) int {
	left, zeros, maxLen := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// LengthOfLongestSubstringKDistinct — длина самой длинной подстроки
// с не более чем k различными символами.
//
// Пример: s="eceba", k=2 -> 3 ("ece").
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	if k <= 0 {
		return 0
	}
	count := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++

		for len(count) > k {
			count[s[left]]--
			if count[s[left]] == 0 {
				delete(count, s[left])
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
