package main

import (
	"fmt"
	"sort"
)

//3. Группировка анаграмм
//Условие: Дан []string из строчных латинских букв.
//Сгруппировать слова, которые являются анаграммами.
//Порядок групп и слов внутри группы не важен.
//Пример:
//["eat","tea","tan","ate","nat","bat"] -> [["eat","tea","ate"], ["tan","nat"], ["bat"]]
//Сигнатура:
//func GroupAnagrams(words []string) [][]string

func GroupAnagrams(words []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range words {
		key := sortWord(word)
		groups[key] = append(groups[key], word)
	}

	result := make([][]string, 0)
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

func sortWord(word string) string {
	// Преобразуем строку в слайс байтов
	bytes := []byte(word)

	// Сортируем слайс байтов по возрастанию
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	// Собираем отсортированные байты обратно в строку
	return string(bytes)
}

func main() {
	words := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(words)

	fmt.Println(GroupAnagrams(words))

}
