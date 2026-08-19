package main

import (
	"fmt"
	"strings"
)

//Условие: Напишите функцию ReverseWords(s string) string,
//которая переворачивает порядок слов в строке,
//разделённых пробелами. Слова состоят из букв.
//Лишние пробелы в начале и конце нужно удалить,
//между словами должен быть ровно один пробел.

func ReverseWords(s string) string {

	words := strings.Fields(s)
	reversed := make([]string, 0, len(words))
	for i := len(words) - 1; i >= 0; i-- {
		reversed = append(reversed, words[i])
	}

	return strings.Join(reversed, " ")
}

func main() {
	s1 := ReverseWords("  hello world  ") // "world hello"
	fmt.Println(s1)
	s2 := ReverseWords("a b c") // "c b a"
	fmt.Println(s2)
}
