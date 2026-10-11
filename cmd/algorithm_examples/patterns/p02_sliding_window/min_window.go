package main

// Паттерн: сжимающееся окно (minimum-size sliding window).
//
// Идея: окно расширяется вправо, пока не выполнит условие. Как только
// условие выполнено — пытаемся сжать слева, обновляя минимум. Затем
// продолжаем расширять вправо.
//
// Отличие от variable_window: там мы ищем МАКСИМУМ и сжимаем только
// при нарушении условия. Здесь ищем МИНИМУМ и сжимаем всегда, когда
// условие выполнено.
//
// Когда применять:
//   - "минимальная подстрока/подмассив, содержащая ...";
//   - "минимальная длина окна с суммой >= target";
//   - "наименьшее окно, удовлетворяющее условию".
//
// Сложность: O(n) время, O(k) память, где k — размер алфавита.

// MinSubArrayLen — минимальная длина подмассива с суммой >= target.
// Если такого подмассива нет — вернуть 0.
//
// Пример: target=7, nums=[2,3,1,2,4,3] -> 2 (подмассив [4,3]).
func MinSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	minLen := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Условие выполнено — пробуем сжать окно.
		for sum >= target {
			windowLen := right - left + 1
			if minLen == 0 || windowLen < minLen {
				minLen = windowLen
			}
			sum -= nums[left]
			left++
		}
	}
	return minLen
}

// MinWindow — минимальная подстрока s, которая содержит ВСЕ символы t
// (с учётом кратности). Если такой подстроки нет — "".
//
// Пример: s="ADOBECODEBANC", t="ABC" -> "BANC".
func MinWindow(s, t string) string {
	if len(s) < len(t) || t == "" {
		return ""
	}

	// need — сколько каких символов требует t.
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	left, formed, required := 0, 0, len(need)
	bestLeft, bestLen := 0, 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		have[ch]++
		// Символ впервые достиг нужного количества.
		if have[ch] == need[ch] {
			formed++
		}

		// Все символы собраны — пробуем сжать окно.
		for formed == required {
			windowLen := right - left + 1
			if bestLen == 0 || windowLen < bestLen {
				bestLeft, bestLen = left, windowLen
			}

			leftCh := s[left]
			have[leftCh]--
			// Потеряли обязательный символ.
			if have[leftCh] < need[leftCh] {
				formed--
			}
			left++
		}
	}

	if bestLen == 0 {
		return ""
	}
	return s[bestLeft : bestLeft+bestLen]
}
