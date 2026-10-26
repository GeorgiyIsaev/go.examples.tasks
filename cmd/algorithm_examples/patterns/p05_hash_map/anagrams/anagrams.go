package main

import "sort"

// Паттерн: хэш-мапа как «отпечаток» (fingerprint).
//
// Идея: чтобы понять, «те же ли это элементы», не сравниваем всё со
// всем, а сводим объект к ключу — отпечатку — и группируем по нему.
// Два приёма:
//
//  1. отпечаток = счётчик символов (частоты);
//  2. отпечаток = отсортированная строка (канонический вид).
//
// Оба дают O(n) на группу вместо O(n²) попарных сравнений.

// IsAnagram — являются ли s и t анаграммами (LeetCode 242).
// Сравниваем по рунам, чтобы работало с Unicode.
//
// Один счётчик: прибавляем буквы из s, вычитаем из t. Если после
// всего прохода где-то неноль — не анаграмма. Длину проверяем сразу.
//
// Пример: "anagram" / "nagaram" -> true; "rat" / "car" -> false.
func IsAnagram(s, t string) bool {
	runes := []rune(s)
	if len(runes) != len([]rune(t)) {
		return false
	}
	count := map[rune]int{}
	for _, r := range runes {
		count[r]++
	}
	for _, r := range t {
		count[r]--
		if count[r] < 0 {
			// Зашли в минус — в t есть лишняя буква.
			return false
		}
	}
	// Длины равны, значит положительных остатков не осталось.
	return true
}

// GroupAnagrams — сгруппировать анаграммы (LeetCode 49).
// Порядок групп и строк внутри группы не важен.
//
// Ключ группы — отсортированная строка ("eat" -> "aet"):
// два слова анаграммы <=> у них одинаковый отсортированный вид.
// Альтернатива — ключ-счётчик вида "a1e1t1".
//
// Пример: ["eat","tea","tan"] -> [["eat","tea"],["tan"]].
func GroupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, w := range strs {
		runes := []rune(w)
		sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
		key := string(runes)
		groups[key] = append(groups[key], w)
	}

	result := make([][]string, 0, len(groups))
	for _, g := range groups {
		result = append(result, g)
	}
	return result
}
