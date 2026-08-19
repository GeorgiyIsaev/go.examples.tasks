package main

import "fmt"

// Напишите функцию LongestUniqueSubstr(s string) string,
// которая возвращает самую длинную подстроку,
// не содержащую повторяющихся символов
// (если несколько — вернуть первую).

func LongestUniqueSubstr(s string) string {
	// Преобразуем строку в срез рун для корректной работы с Unicode
	runes := []rune(s)
	n := len(runes)
	if n == 0 {
		return ""
	}

	// Карта: символ -> его последняя позиция (индекс в runes)
	lastPos := make(map[rune]int)
	left := 0     // левая граница текущего окна
	maxStart := 0 // начало самой длинной найденной подстроки
	maxLen := 0   // длина самой длинной подстроки

	for right := 0; right < n; right++ {
		ch := runes[right]

		// Если символ уже был в текущем окне, сдвигаем левую границу
		if pos, ok := lastPos[ch]; ok && pos >= left {
			left = pos + 1
		}

		// Обновляем позицию текущего символа
		lastPos[ch] = right

		// Текущая длина окна
		curLen := right - left + 1
		// Если нашли более длинную подстроку (строго больше), сохраняем её
		if curLen > maxLen {
			maxLen = curLen
			maxStart = left
		}
	}

	// Извлекаем подстроку из среза рун и возвращаем как строку
	return string(runes[maxStart : maxStart+maxLen])
}

func main() {
	fmt.Println(LongestUniqueSubstr("abcabcbb")) // "abc"
	fmt.Println(LongestUniqueSubstr("bbbbb"))    // "b"
	fmt.Println(LongestUniqueSubstr("pwwkew"))   // "wke"
	fmt.Println(LongestUniqueSubstr(""))         // ""
	fmt.Println(LongestUniqueSubstr("aab"))      // "ab"
}
