package main

import (
	"fmt"
	"sort"
	"strings"
)

//Напишите функцию topWords(text string, n int) []string,
//которая принимает текст и число n,
//а возвращает n самых частых слов в порядке убывания частоты.
//Если слова имеют одинаковую частоту — порядок не важен.

type wordCount struct {
	word  string
	count int
}

func topWords(text string, n int) []string {
	words := strings.Fields(text) //слайс слов
	fmt.Println("words=", words)
	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}

	counts := make([]wordCount, 0, len(freq))
	for word, count := range freq {
		counts = append(counts, wordCount{word, count})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	result := make([]string, 0, n)
	for i := 0; i < n && i < len(counts); i++ {
		result = append(result, counts[i].word)
	}
	return result
}

func main() {
	text := "cat dog cat bird dog cat bird bird bird"
	result := topWords(text, 2) // ["bird", "cat"] (bird — 4, cat — 3, dog — 2)

	fmt.Println("result", result)
}
