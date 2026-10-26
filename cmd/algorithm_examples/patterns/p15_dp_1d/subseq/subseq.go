package main

import "sort"

// Подпоследовательности: базовые DP 1D на «я подхожу/не подхожу».
//
// Ключевое отличие от подмассива (окно, p03): подпоследовательность
// НЕПРЕРЫВНАЯ — можно пропускать элементы, но порядок сохраняется.
// Отсюда «два индекса» в рекуррентности => O(n²), либо более хитрые
// структуры => O(n log n).

// LengthOfLIS — длина самой длинной возрастающей подпоследовательности
// (LeetCode 300).
//
// O(n²)-DP: dp[i] = длина LIS, заканчивающейся элементом nums[i].
// dp[i] = max(dp[j] + 1) по всем j < i с nums[j] < nums[i]; иначе 1.
// Ответ — max по dp.
//
// Пример: [10,9,2,5,3,7] -> 3 ([2,5,7] или [2,3,7]).
// Строгое возрастание: [10,9,2,5,3,7,101,18] -> 4 ([2,3,7,101]).
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	best := 1
	for i := range dp {
		dp[i] = 1 // минимум — сам элемент
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > best {
			best = dp[i]
		}
	}
	return best
}

// LengthOfLISFast — то же за O(n log n) «методом patience sorting»:
// tails[k] — наименьший последний элемент всех возрастающих
// подпоследовательностей длины k+1, что мы уже видели.
// Каждый элемент binary-search-ом встаёт в свой хвост:
// заменяет первый >= x (если такого нет — хвосты растут => LIS длиннее).
//
// tails отсортирован — поэтому работает sort.Search.
func LengthOfLISFast(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		idx := sort.SearchInts(tails, x)
		if idx == len(tails) {
			tails = append(tails, x)
		} else {
			tails[idx] = x
		}
	}
	return len(tails)
}

// IsSubsequence — проверка «s — подпоследовательность t»
// (LeetCode 392, Easy): два указателя — жадный проход.
// Классическая связка с DP: ЛКС длины |s| возвращает |s| целиком.
//
// Пример: s="abc", t="ahbgdc" -> true.
func IsSubsequence(s, t string) bool {
	i := 0
	for j := 0; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}
