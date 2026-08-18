package main

import (
	"fmt"
	"sort"
)

// Напишите функцию groupAnagrams(words []string) [][]string,
// которая принимает список слов и возвращает список списков,
// где каждый внутренний список содержит слова,
// являющиеся анаграммами друг друга (состоят из одних и тех же букв).

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[string][]string)
	for _, word := range words {
		key := sortString(word)
		anagramMap[key] = append(anagramMap[key], word)
	}
	fmt.Println("m=:", anagramMap)

	result := make([][]string, 0, len(anagramMap))
	for _, anagram := range anagramMap {
		result = append(result, anagram)
	}
	return result
}

func main() {
	words := []string{"listen", "silent", "enlist", "inlets", "google", "gogole", "cat", "act", "tac"}
	// Ожидаемый результат: [["listen","silent","enlist","inlets"], ["google","gogole"], ["cat","act","tac"]
	m := groupAnagrams(words)
	fmt.Println("m=:", m)
}
