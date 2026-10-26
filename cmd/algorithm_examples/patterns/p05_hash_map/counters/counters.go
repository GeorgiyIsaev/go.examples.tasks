package main

// Паттерн: хэш-мапа как СЧЁТЧИК.
//
// Идея: вместо вложенных циклов O(n²) один проход, который на лету
// «спрашивает у карты»: сколько раз встречалось / где встречалось.
//
// Базовые операции: m[key]++ (прибавить), m[key] (прочитать с нулём
// по умолчанию), delete(m, key). Амортизированная сложность O(1)
// на операцию => проход по n элементам — O(n) времени, O(n) памяти.
//
// Три приёма, которые здесь разобраны:
//  1. частота элемента (счётчик);
//  2. «элемент встречается больше n/2 раз» (majority);
//  3. топ-K по частоте (map + сортировка по счётчику).

import "sort"

// CharFrequency — сколько раз каждый символ встречается в строке.
// Ключ — руна, чтобы не резать многобайтовые символы (UTF-8).
//
// Пример: "banana" -> {a:3, b:1, n:2}.
func CharFrequency(s string) map[rune]int {
	freq := map[rune]int{}
	for _, r := range s {
		freq[r]++
	}
	return freq
}

// MajorityElement — элемент, встречающийся больше n/2 раз (LeetCode 169).
// По условию такой элемент гарантированно есть.
//
// Счётчик за один проход: как только счётчик превысил n/2 — это и есть
// ответ, дальше можно не считать.
//
// Пример: [2,2,1,1,1,2,2] -> 2.
func MajorityElement(nums []int) int {
	counts := map[int]int{}
	majority := len(nums) / 2
	for _, v := range nums {
		counts[v]++
		if counts[v] > majority {
			return v
		}
	}
	// Невозможно при корректном условии задачи; для надёжности — ноль.
	return 0
}

// TopKFrequent — k наиболее частых элементов (LeetCode 347).
// Порядок ответа не важен.
//
// Шаг 1: map за O(n). Шаг 2: сортируем ОТЛИЧАЮЩИЕСЯ ключи по
// убыванию счётчика — O(m log m), m — число уникальных.
// Второй шаг можно заменить bucket sort (массив корзин по частоте)
// на O(n), но map + sort проще и часто достаточен.
//
// Пример: [1,1,1,2,2,3], k=2 -> [1 2].
func TopKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	for _, v := range nums {
		counts[v]++
	}

	unique := make([]int, 0, len(counts))
	for v := range counts {
		unique = append(unique, v)
	}
	sort.Slice(unique, func(i, j int) bool {
		if counts[unique[i]] != counts[unique[j]] {
			return counts[unique[i]] > counts[unique[j]] // чаще — вперёд
		}
		return unique[i] < unique[j] // детерминизм при равных
	})
	return unique[:k]
}
