package main

import "fmt"

//8. Самая длинная подстрока без повторяющихся символов
//Условие: Дана строка. Вернуть длину самой длинной подстроки,
//в которой все символы уникальны.
//Примеры:
//•	"abcabcbb" → 3 ("abc")
//•	"bbbbb" → 1 ("b")
//•	"pwwkew" → 3 ("wke")
//•	"" → 0
//•	"abcdef" → 6
//Сигнатура:
//func LengthOfLongestSubstring(s string) int

func LengthOfLongestSubstring(s string) int {
	last := make(map[rune]int)
	left := 0
	maxLen := 0
	pos := 0 // позиция текущего символа по порядку

	for _, ch := range []rune(s) {
		// Если символ уже встречался внутри текущего окна,
		// сдвигаем левую границу сразу после него.
		if prev, ok := last[ch]; ok && prev >= left {
			left = prev + 1
		}

		last[ch] = pos

		// Текущая длина окна: от left до pos включительно.
		if curLen := pos - left + 1; curLen > maxLen {
			maxLen = curLen
		}

		pos++
	}

	return maxLen
}

func main() {
	fmt.Println(LengthOfLongestSubstring("abcabcbb"))
	fmt.Println(LengthOfLongestSubstring("bbbbb"))
	fmt.Println(LengthOfLongestSubstring("pwwkew"))
	fmt.Println(LengthOfLongestSubstring(""))
	fmt.Println(LengthOfLongestSubstring("abcdef"))
}
