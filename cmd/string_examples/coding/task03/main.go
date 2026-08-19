package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Напишите функцию Compress(s string) string,
//которая выполняет сжатие повторяющихся символов:
//заменяет последовательности одинаковых символов
//на символ и число повторений.
//Если сжатая строка не короче исходной, вернуть исходную.

func Compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Переводим строку в срез рун для корректной работы с Unicode
	runes := []rune(s)
	n := len(runes)

	var builder strings.Builder
	count := 1
	prev := runes[0]

	for i := 1; i < n; i++ {
		cur := runes[i]
		if cur == prev {
			count++
		} else {
			// Записываем предыдущий символ
			builder.WriteRune(prev)
			if count > 1 {
				builder.WriteString(strconv.Itoa(count))
			}
			// Сбрасываем для нового символа
			prev = cur
			count = 1
		}
	}
	// Добавляем последнюю группу
	builder.WriteRune(prev)
	if count > 1 {
		builder.WriteString(strconv.Itoa(count))
	}

	compressed := builder.String()
	// Если сжатая строка короче исходной, возвращаем её, иначе исходную
	if len(compressed) < len(s) {
		return compressed
	}
	return s
}

func main() {
	fmt.Println(Compress("aaabbc")) // "a3b2c"
	fmt.Println(Compress("abc"))    // "abc"
	fmt.Println(Compress(""))       // ""
	fmt.Println(Compress("aab"))    // "a2b" (длина 3, исходная 3, не короче → возвращаем "aab"? но сжатая "a2b" длина 3, не короче, значит возвращаем исходную "aab") - проверка
	fmt.Println(Compress("aaa"))    // "a3" (длина 2 < 3 → возвращаем "a3")
	fmt.Println(Compress("hello"))  // "hel2o"? нет, "hello" -> h e l2 o? буква l повторяется дважды, сжатая "hel2o" длина 5, исходная 5, не короче → возвращаем "hello"
}
